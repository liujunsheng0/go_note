package shape

import (
	"fmt"
	"math"
)

// 图形
type Shaper interface {
	Name() string
	Area() float64 // 面积
}

// 以下结构体均实现了Shaper接口
// 圆形
type Circle struct {
	radius float64
}

// 正方形
type Square struct {
	side float64
}

// 长方形
type Rectangle struct {
	length, width float64
}

func (s Circle) Area() float64 {
	return math.Pi * s.radius * s.radius
}

func (_ Circle) Name() string {
	return "circle"
}

func (s Circle) String() string {
	return fmt.Sprintf("circle: radius=%v", s.radius)
}

// 接收者为实例
func (s Rectangle) Area() float64 {
	return s.length * s.width
}

// 接收者为实例
func (_ Rectangle) Name() string {
	return "rectangle"
}

func (s Rectangle) String() string {
	return fmt.Sprintf("rectangle: width=%v length=%v", s.width, s.length)
}

// 接收者为实例
func (s Square) Area() float64 {
	return s.side * s.side
}

// 接收者为实例
func (_ Square) Name() string {
	return "square"
}

func (s Square) String() string {
	return fmt.Sprintf("square: side=%v", s.side)
}
