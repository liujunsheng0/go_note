package sync_

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// expunged entry: 已删除的实体
// unexpunged entry: 未删除的实体
type Map struct {
	mu sync.Mutex // 保护 read 和 dirty
	// read 可以并发访问，load时需要持有锁
	// read map是被atomic包托管的，这意味着它本身Load是并发安全的（但是它的Store操作需要锁mu的保护）
	// read map中的entries可以安全地并发更新，但是对于expunged entry(已删除)，在更新前需要经它unexpunge化并存入dirty
	//（这句话，在Store方法的第一种特殊情况中，使用e.unexpungeLocked处有所体现）
	read atomic.Value // readOnly
	// 关于dirty map必须要在锁mu的保护下，进行操作。它仅仅存储 non-expunged entries
	// 如果一个 expunged entries需要存入dirty，需要先进行unexpunged化处理
	// 如果dirty map是nil的，则对dirty map的写入之前，需要先根据read map对dirty map进行浅拷贝初始化
	dirty map[interface{}]*entry
	// 每当读取的是时候，read中不存在，需要去dirty查看，miss自增，到一定程度会触发dirty=>read升级转储
	// 升级完毕之后，dirty置空 &miss清零 &read.amended置false
	misses int
}

type readOnly struct {
	m       map[interface{}]*entry
	amended bool // 标记位，如果为真，则说明dirty中存在新增的key，还没升级转储，不存在于read中
}

// expunged 删除  任意指针，用于标记从Map.dirty 中删除
var expunged = unsafe.Pointer(new(interface{}))

// map中对应的value
// An entry is a slot in the map corresponding to a particular key.
type entry struct {
	// p points to the interface{} value stored for the entry.
	//
	// If p == nil, the entry has been deleted and m.dirty == nil.
	// If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
	// is missing from m.dirty.
	// Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty != nil, in m.dirty[key].

	// entry的p可能的状态：
	// e.p == nil：entry已经被标记删除，不过此时还未经过read=>dirty重塑，如果dirty非nil，可能仍然属于dirty (read中有, dirty中可能有)
	// e.p == expunged：entry已经被标记删除，经过read=>dirty重塑，不属于dirty，仅仅属于read，下一次dirty=>read升级，会被彻底清理(read中有, dirty中无)
	// e.p == 普通指针：此时entry是一个不同的存在状态，属于read，如果dirty非nil，也属于dirty

	// An entry can be deleted by atomic replacement with nil: when m.dirty is next created,
	// it will atomically replace nil with expunged and leave m.dirty[key] unset.
	//
	// An entry's associated value can be updated by atomic replacement, provided
	// p != expunged. If p == expunged, an entry's associated value can be updated
	// only after first setting m.dirty[key] = e so that lookups using the dirty
	// map find the entry.
	p unsafe.Pointer // *interface{}
}

func (e *entry) load() (value interface{}, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == nil || p == expunged {
		return nil, false
	}
	return *(*interface{})(p), true
}

// tryStore stores a value if the entry has not been expunged. If the entry is expunged, tryStore returns false and leaves the entry unchanged.
func (e *entry) tryStore(i *interface{}) bool {
	for {
		// 乐观锁
		// 如果p是expunged就不可以了set了
		// 因为expunged状态是read独有的，这种情况下说明这个key已经删除（并且发生过了read=>dirty重塑过）了,此时要新增只能在dirty中，不能在read中
		p := atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
		// 乐观锁, 比较值， 替换
		// 如果非expunged，则说明是normal的entry或者nil的entry，可以直接替换
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
	}
}

// storeLocked unconditionally stores a value to the entry. The entry must be known not to be expunged.
// 存储值到entry, entry必须是未删除的
func (e *entry) storeLocked(i *interface{}) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}

// 共享指针, read中有, dirty中可能有, 如果值发生变化, 这样 read 和 dirty 都能看到这个变化
// 未被删除
// unexpungeLocked ensures that the entry is not marked as expunged.
// If the entry was previously expunged, it must be added to the dirty map before m.mu is unlocked.
func (e *entry) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}

// 已删除 尝试软删除
func (e *entry) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == expunged
}

func (e *entry) delete() (hadValue bool) {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return true
		}
	}
}

func newEntry(i interface{}) *entry {
	return &entry{p: unsafe.Pointer(&i)}
}

// 返回对应key的值, ok: 是否在map中
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	// 如果readonly没找到，且dirty包含了read没有的key，则尝试去dirty里面找
	if !ok && read.amended {
		m.mu.Lock()
		// Avoid reporting a spurious(虚假的) miss if m.dirty got promoted(新增) while we were blocked on m.mu.
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			// 记录miss次数，并在满足阈值后，触发dirty=>read
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return nil, false
	}
	return e.load()
}

func (m *Map) Store(key, value interface{}) {
	// 对应的key值存在, 替换key对应的val
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}
	// 如果readonly里面不存在key或者是对应的key是被擦除掉了的，则继续。。。
	m.mu.Lock()
	// 锁的惯用模式：再次检查readonly，防止在上锁前的时间缝隙出现存储
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok { // read中存在
		// 已删除
		if e.unexpungeLocked() {
			// The entry was previously expunged, which implies that there is a non-nil dirty map and this entry is not in it.
			m.dirty[key] = e
		}
		e.storeLocked(&value)
	} else if e, ok := m.dirty[key]; ok { // dirty中存在, 直接更改值即可, read和dirty指向的是同一个对象
		e.storeLocked(&value)
	} else {
		// dirty里面也不存在（或者dirty为nil），则应该先设置在ditry里面
		// 此时要检查read.amended，如果为假（标识dirty中没有自己独有的key or 两者均是初始化状态）
		// 此时要在dirty里面设置新的key，需要确保dirty是初始化的且需要设置amended为true（表示自此dirty多出了一些独有key）
		if !read.amended {
			// We're adding the first new key to the dirty map.
			// Make sure it is allocated and mark the read-only map as incomplete.
			m.dirtyLocked()
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}
	m.mu.Unlock()
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	// Avoid locking if it's a clean hit.
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		actual, loaded, ok := e.tryLoadOrStore(value)
		if ok {
			return actual, loaded
		}
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		actual, loaded, _ = e.tryLoadOrStore(value)
	} else if e, ok := m.dirty[key]; ok {
		actual, loaded, _ = e.tryLoadOrStore(value)
		m.missLocked()
	} else {
		if !read.amended {
			// We're adding the first new key to the dirty map.
			// Make sure it is allocated and mark the read-only map as incomplete.
			m.dirtyLocked()
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
		actual, loaded = value, false
	}
	m.mu.Unlock()

	return actual, loaded
}

// tryLoadOrStore atomically loads or stores a value if the entry is not
// expunged.
//
// If the entry is expunged, tryLoadOrStore leaves the entry unchanged and
// returns with ok==false.
func (e *entry) tryLoadOrStore(i interface{}) (actual interface{}, loaded, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == expunged {
		return nil, false, false
	}
	if p != nil {
		return *(*interface{})(p), true, true
	}

	// Copy the interface after the first load to make this method more amenable
	// to escape analysis: if we hit the "load" path or the entry is expunged, we
	// shouldn't bother heap-allocating.
	ic := i
	for {
		if atomic.CompareAndSwapPointer(&e.p, nil, unsafe.Pointer(&ic)) {
			return i, false, true
		}
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return nil, false, false
		}
		if p != nil {
			return *(*interface{})(p), true, true
		}
	}
}

// Delete deletes the value for a key.
func (m *Map) Delete(key interface{}) {
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			delete(m.dirty, key)
		}
		m.mu.Unlock()
	}
	if ok {
		e.delete()
	}
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *Map) Range(f func(key, value interface{}) bool) {
	// We need to be able to iterate over all of the keys that were already
	// present at the start of the call to Range.
	// If read.amended is false, then read.m satisfies that property without
	// requiring us to hold m.mu for a long time.
	read, _ := m.read.Load().(readOnly)
	if read.amended {
		// m.dirty contains keys not in read.m. Fortunately, Range is already O(N)
		// (assuming the caller does not break out early), so a call to Range
		// amortizes an entire copy of the map: we can promote the dirty copy
		// immediately!
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		if read.amended {
			read = readOnly{m: m.dirty}
			m.read.Store(read)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}

	for k, e := range read.m {
		v, ok := e.load()
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

// 重新加载
func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	// 直接用dirty覆盖到了read上（那也就是意味着dirty的值是必然是read的父集合，当然这不包括read中的expunged entry）
	// 同时read.amended再次变成false
	m.read.Store(readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

// 新增值&dirty为nil, 将未被删除的entry read->dirty
func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}
	read, _ := m.read.Load().(readOnly)
	m.dirty = make(map[interface{}]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() { // 非nil、expunged的实体，即未被删除的entry
			m.dirty[k] = e
		}
	}
}
