package algorithm

import (
	"container/list"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 最近常使用缓存
// add 存在? 不存在?
// get
// 删除不使用
type KV struct {
	k interface{}
	v interface{}
}

type LRU struct {
	mutex sync.Mutex
	l     *list.List
	kv    map[interface{}]*list.Element
	cap   int
}

func (s *LRU) Put(k, v interface{}) {
	if s == nil {
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 已存在
	el, found := s.kv[k]
	if found {
		el.Value = v
		s.l.MoveToFront(el)
		return
	}
	var kv *KV
	if len(s.kv) >= s.cap {
		kv = s.removeBack()
		kv.k = k
		kv.v = v
	} else {
		kv = &KV{
			k: k,
			v: v,
		}
	}
	el = s.l.PushFront(kv)
	s.kv[k] = el
}

func (s *LRU) Get(k interface{}) (bool, interface{}) {
	if s == nil {
		return false, nil
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	el, found := s.kv[k]
	if !found {
		return false, nil
	}
	s.l.MoveToFront(el)
	return true, el.Value.(*KV).v
}

func (s *LRU) removeBack() *KV {
	if s == nil {
		return nil
	}
	el := s.l.Back()
	if el == nil {
		return nil
	}
	s.l.Remove(el)
	kv := el.Value.(*KV)
	delete(s.kv, kv.k)
	return kv
}

func NewLRUCache(cap int) *LRU {
	if cap < 1 {
		cap = 1
	}
	cache := LRU{
		mutex: sync.Mutex{},
		l:     nil,
		kv:    nil,
		cap:   cap,
	}
	cache.l = list.New()
	cache.kv = make(map[interface{}]*list.Element, cap)
	return &cache
}

func TestLRU(t *testing.T) {
	ar := assert.New(t)
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	_, v := lru.Get(1)
	ar.Equal(1, v)
	lru.Put(3, 3) // 该操作会使得密钥 2 作废
	found, _ := lru.Get(2)
	ar.True(!found)
	lru.Put(4, 4) // 该操作会使得密钥 1 作废
	found, _ = lru.Get(1)
	ar.True(!found)
	_, v = lru.Get(3)
	ar.Equal(3, v)
	_, v = lru.Get(4)
	ar.Equal(4, v)
}
