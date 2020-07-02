package reflect

import (
	"fmt"
	"runtime"
	"testing"
)

func Recover(t *testing.T) {
	if err := recover(); err != nil {
		buf := make([]byte, 1<<18)
		n := runtime.Stack(buf, false)
		if nil == t {
			fmt.Printf("%v, STACK: %s", err, buf[0:n])
		} else {
			t.Logf("%v, STACK: %s", err, buf[0:n])

		}
	}
}
