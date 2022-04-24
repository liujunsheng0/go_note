package channel

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 通道, 实际上是类型化消息的队列, 使数据得以传输, 它是先进先出的结构所以可以保证发送给他们的元素的顺序
// var name chan datatype
func TestInit(t *testing.T) {
	var ch1 chan int      // 未初始化的通道的值是 nil
	ch2 := make(chan int) // 引用类型, 使用make来分配内存

	assert.Equal(t, true, reflect.ValueOf(ch1).IsNil())
	assert.Equal(t, false, reflect.ValueOf(ch2).IsNil())
}

// 通信操作符 <-, 信息按照箭头的方向流动
// 发送  流向通道    ch  <- data
// 接收  从通道流出  data <- ch
// 注意  打印的日志并不能表明通道的发送和接收顺序, 因为打印日志和通道实际发生读写的时间延迟会导致和真实发生的顺序不同
func TestCommunicate(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("send  recv")
	send := func(ch chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(" ", i)
			ch <- i
			time.Sleep(1 * time.Second) // 保证打印顺序和发送, 接收顺序一致
		}
		wg.Done()
	}
	recv := func(ch chan int) {
		num := 0
		for {
			select {
			case i := <-ch:
				assert.Equal(t, num, i)
				num++
				fmt.Println("      ", i)
			case <-time.After(time.Second * 2):
				wg.Done()
				return

			}
		}
	}
	ch := make(chan int, 0)

	go send(ch)
	go recv(ch)
	wg.Wait()
}

// 容量(capacity)代表Channel容纳的最多的元素的数量. 如果没有设置容量, 或者容量设置为0, 说明Channel没有缓存,
// 只有sender和receiver都准备好了后它们的通讯才会发生. 如果设置了缓存, 就有可能不发生阻塞.
// 只有buffer满了后sender才会阻塞, 只有缓存空了后receiver才会阻塞.
// 一个nil channel不会通信
// len(channel): 查看缓存队列中排队个数
func TestCapacity(t *testing.T) {
	capacity := 2
	assert.Equal(t, cap(make(chan int)), 0) // 默认容量为0
	assert.Equal(t, cap(make(chan int, capacity)), capacity)

	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("send  recv  len(ch)")
	send := func(ch chan int) {
		for i := 1; i < 9; i++ {
			// 容量为10, 如果缓存满了以后阻塞, 有位置后, 再放入缓存队列中
			ch <- i
			fmt.Println(" ", i, "         ", len(ch))

		}
		wg.Done()
	}
	recv := func(ch chan int) {
		for {
			time.Sleep(2 * time.Second)
			select {
			case i := <-ch:
				fmt.Println("      ", i, "    ", len(ch))
			case <-time.After(time.Second):
				wg.Done()
				return

			}
		}
	}
	ch := make(chan int, 3)
	go send(ch)
	go recv(ch)
	wg.Wait()
}

// 给 已经关闭的通道 发送或者再次关闭都会导致运行时的panic
// close(chan), 虽然chan已经关闭, 但是仍可以从chan中取出在关闭前缓存的元素, 如果缓存中没有元素, chan会继续返回指定类型的默认值
func TestClose(t *testing.T) {
	sendToCloseChan := func() {
		defer func() {
			t.Log(recover()) // send on closed channel
		}()
		ch := make(chan int)
		close(ch)
		ch <- 1
	}
	sendToCloseChan()

	recvFromCloseChan := func() {
		ch := make(chan int)
		close(ch)
		for times := 0; times < 10; {
			select {
			// v, ok := <- channel
			// 未关闭chan  阻塞等待值  v=接收的值    ok=true
			// 关闭chan    不阻塞     v=类型默认值  ok=false
			case v, ok := <-ch:
				if !ok {
					times++
				}
				t.Log("recv", v, "times", times)
			case <-time.After(time.Millisecond):
				times++
			}
		}
	}
	recvFromCloseChan()

	capacity := 5
	ch := make(chan int, capacity)
	recv := make([]int, 0)
	for i := 1; i < capacity+1; i++ {
		ch <- i
	}
	t.Log("isClosed:", IsClosed(ch))
	close(ch)
	t.Log("isClosed:", IsClosed(ch))
	// 虽然ch已经关闭, 但是仍可以从ch中取出在关闭前缓存的元素
	for v := range ch {
		recv = append(recv, v)
	}
	t.Log("recv", recv)
}

// 判断通道是否关闭
func TestIsClosed(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	close(ch2)

	// v, ok := <- channel
	// 未关闭chan  阻塞等待值  v=接收的值    ok=true
	// 关闭chan    不阻塞     v=类型默认值  ok=false
	v1, isOpen1 := <-ch1
	v2, isOpen2 := <-ch2

	_, isOpen3, _ := Receive(ch1)
	_, isOpen4, _ := Receive(ch2)

	t.Log(fmt.Sprintf("v1:%v  isOpen:%-5v isOpen:%-5v isClosed:%v", v1, isOpen1, isOpen3, IsClosed(ch1)))
	t.Log(fmt.Sprintf("v2:%v  isOpen:%-5v isOpen:%-5v isClosed:%v", v2, isOpen2, isOpen4, IsClosed(ch2)))
}

// for 循环的 range 语句可以用在通道 ch 上, 便可以从通道中获取值, 它从指定通道中读取数据直到*通道关闭*(写入和读取完成后, 结束for循环)
func TestFor(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Duration(i%3) * time.Second)
		}
	}()
	fmt.Printf("recv     time\n")
	for v := range ch {
		fmt.Printf(" %-2v    %v\n", v, Now())
	}
	fmt.Println("---  finish  ---")
}

// 通道的方向
// 通道类型可以用注解来表示它只发送或者只接收, 默认创建的channel是双向的
// var send_only chan <- int          channel can only send data
// var recv_only <- chan int          channel can only recv data
// recv_only无法关闭, 因为关闭通道是发送者用来表示不再给通道发送值了, 所以对只接收通道是没有意义的
func TestDirectionOfChannel(t *testing.T) {
	sendOnly := func(in chan<- int) {
		// <- in      语法错误
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}
	recvOnly := func(out <-chan int) {
		// close(ch)  语法错误
		i := 0
		for v := range out {
			assert.Equal(t, i, v)
			i++
		}
	}
	ch := make(chan int)
	go sendOnly(ch)
	recvOnly(ch)
}

// select 通信开关
// 在任何一个 case 中执行 break 或者 return, select 就结束了
// 使用 select 切换协程, 从不同并发执行的协程中获取值可以通过关键字select完成

// select 做的是, 选择处理多个通信情况中的一个
//    如果都阻塞了, 会等待直到其中一个可以处理
//	  如果多个可以处理, 随机选择一个
//    如果没有通道操作可以处理并且写了 default 语句, 它就会执行default语句, default永远是可运行的
//    如果多个case没有能处理的, select会一直阻塞
//    关闭的channel, 直接执行, 多个关闭的channel时, 在前面的先执行
func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			ch1 <- i
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			ch2 <- fmt.Sprintf("%v", i)
		}
	}()

	fmt.Println("ch1  ch2")
	ok := true
	for ok {
		select {
		case v := <-ch1:
			fmt.Println("", v)
		case v := <-ch2:
			fmt.Println("     ", v)
		case <-time.After(3 * time.Second):
			close(ch1)
			close(ch2)
			ok = false
		}
	}

	fmt.Println("=======")
	select {
	case v := <-ch1:
		fmt.Println("ch1 default", v)
	case v := <-ch2:
		fmt.Println("ch2 default", v)
	default:
		fmt.Println("default")

	}
}

// channel 倒序打印1,2,3,4,5,6
func TestFunc(t *testing.T) {
	n := 10
	ch := make(chan int, 10)

	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)

	for i := range ch {
		defer func(i int) { fmt.Println(i) }(i)
	}
}

// channel 倒序打印1,2,3,4,5,6
func TestFunc1(t *testing.T) {
	n := 10
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	close(ch1)
	close(ch2)

	for i := range ch1 {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	for i := range ch2 {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}
