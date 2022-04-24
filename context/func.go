package context

import "sync/atomic"

// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}

// 当父context被取消时, 取消子context
// propagateCancel arranges(安排) for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
	done := parent.Done()
	if done == nil {
		return // parent is never canceled
	}

	// 不会阻塞, channel没有数据时, 直接走default部分
	select {
	case <-done:
		// done -> channel未close
		// parent is already canceled
		child.cancel(false, parent.Err())
		return
	default:
	}

	// done -> channel 已经close
	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent has already been canceled
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		atomic.AddInt32(&goroutines, +1)
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}

// 查找parent中的cancelCtx, cancelCtx := parent.Value(&cancelCtxKey)
// parentCancelCtx returns the underlying(潜在的) *cancelCtx for parent.
// It does this by looking up parent.Value(&cancelCtxKey) to find the *cancelCtx and then checking whether parent.Done() matches that *cancelCtx.
//(If not, the *cancelCtx has been wrapped in a custom implementation providing a different done channel, in which case we should not bypass it.)
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	done := parent.Done()
	if done == closedchan || done == nil {
		return nil, false
	}
	p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
	if !ok {
		return nil, false
	}
	p.mu.Lock()
	ok = p.done == done
	p.mu.Unlock()
	if !ok {
		return nil, false
	}
	return p, true
}

// 从父context中, 移出子context
// removeChild removes a context from its parent.
func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}
