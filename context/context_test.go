package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/
// https://zhuanlan.zhihu.com/p/68792989
func TestContext_WithValue(t *testing.T) {
	ctx := context.TODO()
	ctxA := context.WithValue(ctx, "a", 1)
	ctxB := context.WithValue(ctxA, "b", 2)
	ctxC := context.WithValue(ctxB, "c", 3)
	fmt.Println(ctxC.Value("a"))
	fmt.Println(ctxC.Value("b"))
	fmt.Println(ctxC.Value("c"))
}

func TestContext_WithCancel(t *testing.T) {
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, sleep int) {
			select {
			case <-ctx.Done():
				t.Logf("ctxDone:%vs", sleep)
				return
			case <-time.After(time.Duration(sleep) * time.Second):
				t.Logf("sleep:%vs", sleep)
			}
		}(cancelCtx, i)
	}
	// sleep 1s
	// sleep 2s
	time.Sleep(2 * time.Second)
	cancelFunc()
	time.Sleep(time.Second)
	t.Log("finish")
}
