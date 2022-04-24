package context

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/
// https://zhuanlan.zhihu.com/p/68792989
func TestContext_WithValue(t *testing.T) {
	ar := assert.New(t)
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "a", "b")
	s, _ := ctx.Value("a").(string)
	ar.Equal(nil, ctx.Value("a"))
	ar.Equal("", s)
	ar.Equal("b", ctx1.Value("a"))

	t.Log(ctx1.Deadline())
	t.Log(ctx1.Done())
	t.Log(ctx1.Err())
}

func TestContext_ChannelClose(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(sleep int) {
			select {
			case <-ch:
				t.Logf("ch sleep:%v", sleep)
				return
			case <-time.After(time.Duration(sleep) * time.Second):
				t.Logf("sleep:%vs", sleep)
			}
		}(i)
	}
	t.Log("runtime.NumGoroutine()", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	close(ch)
	t.Log("runtime.NumGoroutine()", runtime.NumGoroutine())
	t.Log("finish")
}

func TestContext_WithCancel(t *testing.T) {
	ar := assert.New(t)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, sleep int) {
			select {
			case <-ctx.Done():
				t.Logf("sleep:%vs done", sleep)
				return
			case <-time.After(time.Duration(sleep) * time.Second):
				t.Logf("sleep:%vs", sleep)
			}
		}(cancelCtx, i)
	}
	t.Log("runtime.NumGoroutine()", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	ar.Equal(nil, cancelCtx.Err())
	cancelFunc()
	ar.Equal(context.Canceled, cancelCtx.Err())

	time.Sleep(time.Second)
	t.Log("runtime.NumGoroutine()", runtime.NumGoroutine())
	t.Log("finish")
}
