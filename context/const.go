package context

// goroutines counts the number of goroutines ever created; for testing.
var goroutines int32

// &cancelCtxKey is the key that a cancelCtx returns itself for.
var cancelCtxKey int

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}
