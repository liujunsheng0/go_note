package ratelimit

import (
	"testing"
	"time"

	"github.com/juju/ratelimit"
	"github.com/stretchr/testify/assert"
)

/*
https://chai2010.cn/advanced-go-programming-book/ch5-web/ch5-06-ratelimit.html
流量限制的手段最常见的:漏桶、令牌桶两种
	漏桶是指有一个一直装满了水的桶, 每过固定的一段时间即向外漏一滴水. 如果你接到了这滴水, 那么你就可以继续服务请求, 如果没有接到, 那么就需要等待下一滴水.
	令牌桶则是指匀速向桶中添加令牌, 服务请求时需要从桶中获取令牌, 令牌的数目可以按照需要消耗的资源进行相应的调整. 如果没有令牌, 可以选择等待, 或者放弃.

这两种方法看起来很像, 不过还是有区别的.
	漏桶流出的速率固定;
	令牌桶只要在桶中有令牌, 那就可以拿. 也就是说令牌桶是允许一定程度的并发的, 比如同一个时刻, 有100个用户请求, 只要令牌桶中有100个令牌, 那么这100个请求全都会放过去. 令牌桶在桶中没有令牌的情况下也会退化为漏桶模型.

桶中的令牌数可能是负数
桶初始是满的
func NewBucket(fillInterval time.Duration, capacity int64) *Bucket
默认的令牌桶, fillInterval指每过多长时间向桶里放一个令牌, capacity是桶的容量, 超过桶容量的部分会被直接丢弃. 桶初始是满的.

func NewBucketWithQuantum(fillInterval time.Duration, capacity, quantum int64) *Bucket
和普通的NewBucket()的区别是, 每次向桶中放令牌时, 是放quantum个令牌, 而不是一个令牌.

func NewBucketWithRate(rate float64, capacity int64) *Bucket
会按照提供的比例, 每秒钟填充令牌数. 例如capacity是100, 而rate是0.1, 那么每秒会填充10个令牌.
*/

var (
	fillInterval = time.Millisecond * 200                             // 200毫秒
	capacity     = 12                                                 // 桶容量
	bucket       = ratelimit.NewBucket(fillInterval, int64(capacity)) // 桶初始是满的, 每过fillInterval时间, 往桶中放入一个令牌
)

// 返回拿到count个令牌需要等待的时间, 桶中的令牌已经消费, so可用的令牌数可能为负值
func TestTake(t *testing.T) {
	ar := assert.New(t)
	var count int64 = 10
	ar.Equal(time.Duration(0), bucket.Take(count))
	ar.True(int(bucket.Take(count)) > 0)
	ar.Equal(-8, int(bucket.Available()))
	time.Sleep(fillInterval * 5)
	ar.Equal(-3, int(bucket.Available()))
	time.Sleep(fillInterval * 5)
	ar.Equal(2, int(bucket.Available()))
}

// 消费了令牌
// TakeAvailable(count): 获取可用的令牌, 不足count时, 返回桶中可用的令牌
// 获取后, 桶中的令牌数减去返回的令牌
func TestTakeAvailable(t *testing.T) {
	ar := assert.New(t)
	var count int64 = 5
	ar.Equal(5, int(bucket.TakeAvailable(count)))
	ar.Equal(5, int(bucket.TakeAvailable(count)))
	ar.Equal(2, int(bucket.TakeAvailable(count)))
	ar.Equal(0, int(bucket.TakeAvailable(count)))
	ar.Equal(0, int(bucket.TakeAvailable(count)))
	time.Sleep(time.Second)
	ar.Equal(5, int(bucket.TakeAvailable(count)))
}

// bucket.TakeMaxDuration -> 等待的时间, 等待的时间是否在waitTime之内(bool)
// 返回为False时, 桶中的令牌不会消费
func TestTakeMaxWait(t *testing.T) {
	ar := assert.New(t)
	waitTime, ok := bucket.TakeMaxDuration(10, 0)
	ar.True(ok)
	ar.True(int(waitTime) == 0)
	ar.Equal(2, int(bucket.Available()))
	waitTime, ok = bucket.TakeMaxDuration(5, fillInterval)
	ar.False(ok)
	ar.True(int(waitTime) == 0)
	ar.Equal(2, int(bucket.Available()))
	waitTime, ok = bucket.TakeMaxDuration(5, fillInterval*4)
	ar.True(ok)
	ar.True(waitTime < 3*fillInterval)
	ar.Equal(-3, int(bucket.Available()))
}

// 消费了令牌, 不足时, wait中阻塞, 直至满足条件
// Wait takes count tokens from the bucket, waiting until they are available.
func TestWait(t *testing.T) {
	ar := assert.New(t)
	now := time.Now()
	bucket.Wait(10)
	ar.WithinDuration(time.Now(), now, time.Millisecond)
	ar.Equal(2, int(bucket.Available()))
	bucket.Wait(10)
	ar.WithinDuration(time.Now(), now, fillInterval*9) // fillInterval * 8 + 程序运行时间
	ar.Equal(0, int(bucket.Available()))
}

// WaitMaxDuration
// 返回false 不消费桶中的令牌
// 返回true  消费了桶中的令牌
func TestWaitMaxDuration(t *testing.T) {
	ar := assert.New(t)
	now := time.Now()
	ok := bucket.WaitMaxDuration(10, 0)
	ar.True(ok)
	ar.WithinDuration(time.Now(), now, time.Millisecond)
	ar.Equal(2, int(bucket.Available()))

	ok = bucket.WaitMaxDuration(10, 0)
	ar.False(ok)
	ar.Equal(2, int(bucket.Available()))

	now = time.Now()
	ok = bucket.WaitMaxDuration(10, fillInterval*9)
	ar.True(ok)
	ar.Equal(0, int(bucket.Available()))
	ar.WithinDuration(time.Now(), now, fillInterval*9)
}
