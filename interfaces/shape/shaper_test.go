package shape

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

var (
	s  Shaper
	c  = Circle{1}
	r  = Rectangle{1, 2}
	sq = Square{1}
	// 实现了Shaper接口的类型, 可以赋值给Shaper的实例
	// shapes中的变量指向了实现接口类型实例的引用
	// 接口变量里包含了接收者实例的值和指向对应方法表的指针
	// 多态: 同一种类型在不同的实例上似乎表现出不同的行为
	shapes = []Shaper{s, c, r, sq, &c, &r, &sq}
)

// 接口在标准库中的应用
func TestUsage(t *testing.T) {
	// r 右边的类型都实现了 Read() 方法, 并且有相同的方法签名, r的静态类型是 io.Reader
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("doc.md")
	r = bufio.NewReader(f)
}

// Shaper和实现其接口的类型
func TestShaper(t *testing.T) {
	t.Log("name        area   type               string")
	for _, i := range shapes {
		// 接口变量里包含了接收者实例的值和指向对应方法表的指针
		name := "nil"
		area := 0.0
		if i != nil {
			area = i.Area()
			name = i.Name()
		}
		t.Log(fmt.Sprintf("%-10v  %0.2f   %-17T  %v", name, area, i, i))
	}
}

// 接口实例和类型断言, 判断接口实例具体属于哪个类型
func TestTypeAssert(t *testing.T) {
	for i, v := range shapes {
		switch v := v.(type) {
		// 少了*Square
		case Circle, *Circle, Rectangle, *Rectangle, Square:
			t.Log(fmt.Sprintf("%v  %-18T-> %v", i, v, v))
		case nil:
			t.Log(fmt.Sprintf("%v  -> nil", i))
		default:
			t.Log(fmt.Sprintf("%v  %-18v-> %v", i, "not expected type", v))
		}
	}

	t.Log()
	for id, i := range shapes {
		v, ok := i.(Circle)
		if ok {
			t.Log(fmt.Sprintf("%v -> %v", id, v))
		} else {
			t.Log(fmt.Sprintf("%v -> not circle", id))
		}
	}
}

// 假定 v 是一个值, 想测试它是否实现了某个方法, 可用接口的方式来判断
func TestMethodExist(t *testing.T) {
	// 方法1, 接口
	type namer interface {
		Name() string
	}

	f := func(i interface{}) {
		v, ok := interface{}(i).(namer)
		if ok {
			t.Log("implements Name(),    Name:", v.Name())
		} else {
			t.Log("not implements Name() type:", fmt.Sprintf("%T", i))
		}
		// 方法2: 反射
		if ok = false; i != nil {
			_, ok = reflect.TypeOf(i).MethodByName("Name")
		}
		if ok {
			t.Log("implements Name(),    Name:", v.Name())
		} else {
			t.Log("not implements Name() type:", fmt.Sprintf("%T", i))
		}
	}
	f(s)
	f(c)
	f(r)
	f(sq)
	f(1)

}
