package recover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// recover被用于从 panic 或 错误场景中恢复, 让程序可以重新获得控制权, 停止终止过程进而恢复正常执行
// recover 只能在defer修饰的函数中使用, 用于取得panic调用中传递过来的错误值, 如果是正常执行, 调用recover会返回nil, 且没有其它效果
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.NotNil(t, err)
		}
	}()

	panic("test recover")
}
