package learn

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 造成死锁
// func Deadlock_demo() {
// 	// 无缓冲的通道
// 	ch := make(chan int)
// 	// 若只有发送没有接收 程序会死锁
// 	ch <- 10
// 	fmt.Println("发送成功")
// }

// 同步化通信
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func Sync_channel() {
	ch := make(chan int)
	go recv(ch) // 创建一个 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

// 判断channel里是否还有值可以接收
// 不优雅版
func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

// 优雅版
func f3(ch chan int) {
	// range自带判断channel是否还有值
	for v := range ch {
		fmt.Println(v)
	}
}

func No_elegant_mod() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}

// 单向通道

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Producer2 返回一个接收通道
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func Test_singleway_channel() {
	// 参数为不限制方向的通道
	ch := Producer()

	res := Consumer(ch)
	fmt.Println(res) // 25

	// 参数为限制方向的通道
	ch2 := Producer2()

	res2 := Consumer2(ch2)
	fmt.Println(res2) // 25
}

// 只输出数列中的奇数
func Just_jishu() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
		// 原理为select只能走当前满足条件的一种状态
		// 即使有多个，也只会随机走一个
	}
}

// 示例1
// 各位读者可以查看以下示例代码，尝试找出其中存在的问题。

// demo1 通道误用导致的bug
func demo1() {
	wg2 := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg2.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				// TODO 这里需要判断是否正常取出值
				// 不然一直取零值
				// 这里可增加一个接收值来判断，或改写成select
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task)
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
}

// 示例2
// 各位读者阅读下方代码片段，尝试找出其中存在的问题。

// demo2 通道误用导致的bug
func demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	// 这里程序不会命中case的接收语句
	// 而这是个无缓冲通道
	// 只有发送没有接收程序会死锁
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}

// 没有加锁 产生内存竞争
var (
	x int64

	wg3 sync.WaitGroup // 等待组
)

// add 对全局变量x执行5000次加1操作
func add1() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg3.Done()
}

func Memory_compete() {
	wg3.Add(2)

	go add1()
	go add1()

	wg3.Wait()
	fmt.Println(x)
}

// sync.Mutex 加锁后
var (
	x1 int64

	wg4 sync.WaitGroup // 等待组

	mu1 sync.Mutex // 互斥锁
)

// add 对全局变量x执行5000次加1操作
func add2() {
	for i := 0; i < 5000; i++ {
		mu1.Lock() // 修改x前加锁
		x1 = x1 + 1
		mu1.Unlock() // 改完解锁
	}
	wg4.Done()
}

func Sync_mutex() {
	wg4.Add(2)

	go add2()
	go add2()

	wg4.Wait()
	fmt.Println(x1)
}

var (
	// x       int64
	// wg      sync.WaitGroup
	// mutex   sync.Mutex   //互斥锁
	rwMutex sync.RWMutex //读写锁
)

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)

}

func Tesst_time_spend() {
	// 使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 1000) // x:10 cost:1.466500951s

	// 使用读写互斥锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms
}

// 普通map并不是线程安全的
var mp1 = make(map[string]int)

func get(key string) int {
	return mp1[key]
}

func set(key string, value int) {
	mp1[key] = value
}

func Map_not_safe() {
	// 报错：fatal error: concurrent map writes
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 并发安全的map
var m2 = sync.Map{}

func Map_safe() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 存储key-value
			value, _ := m2.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//

type Counter interface {
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func Test_each_lockway() {
	c1 := CommonCounter{} // 非并发安全
	test(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}
