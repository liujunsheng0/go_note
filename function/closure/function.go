package closure

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// 类似于Python中的装饰器, 但是仅适用于参数为 ...int, 返回值为int的函数
// 计算函数运行时间
func CalculateRunTime(f func(args ...int) int, wg *sync.WaitGroup, t *testing.T) func(args ...int) int {
	return func(args ...int) int {
		defer wg.Done()
		start := time.Now()
		res := f(args...)
		end := time.Now()
		delta := end.Sub(start)
		t.Log(fmt.Sprintf("%-12v= %-3v cost: %v - %v=%v",
			fmt.Sprintf("%v(%v)", strings.Split(FunctionName(f), ".")[1], args),
			res,
			end.Format("04:05"),
			start.Format("04:05"),
			delta.Seconds()))
		return res
	}
}

// 切片的和
func Sum(args ...int) int {
	time.Sleep(time.Second + time.Duration(len(args))*time.Millisecond)
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

// 切片的乘积
func Multi(args ...int) int {
	ans := 1
	for _, v := range args {
		ans *= v
	}
	return ans
}

// 幂
func Pow(args ...int) int {
	for i := 0; i < 100+rand.Intn(100); i++ {
		time.Sleep(10 * time.Millisecond)
	}
	return int(math.Pow(float64(args[0]), float64(args[1])))
}
