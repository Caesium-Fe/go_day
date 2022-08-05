package day10

import (
	"fmt"
	"sync"
	"time"
)

// 题目
func Select_deadlock_demo1() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <-bar:
		default:
			println("default")
		}
	}()
	wg.Wait()
} // 并不会输出default，而是发生死锁

// 另一个例子
func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			// select特性，会先把<-input1,<-input2进行求值
			// 但由于只能执行一个case语句
			// 最终只会从两个值中选一个存到ch中
			// 最后就会发生只存了一半的值进ch
			// 而最后取ch值时，是按照完整的取
			// 就会正常输出五次后，报死锁的错
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
			// 如果改为这样就一切正常：
			// select {
			// case t := <-input1:
			// 	ch <- t
			// case t := <-input2:
			// 	ch <- t
			// }
		}
	}()
	return ch
}

func Select_deadlock_demo2() {
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}

// 第三个例子
func Select_deadlock_demo3() {
	ch := make(chan int, 10)

	go func() {
		var i = 1
		for {
			i++
			ch <- i
		}
	}()

	// 原因是在for+select+time处理的情境下
	// 如果 case <- time.After(30 * time.Second) 这样创建定时任务
	// 会因为for循环每次调用select都会重新初始化一个全新的计时器（Timer）
	// 从而导致一直在不停创建任务，而任务也不能被释放，最终造成内存溢出。
	// 解决办法：
	// 在for循环前可以提前声明定时任务，并保存在变量中，case语句直接调用就好。
	for {
		select {
		case x := <-ch:
			println(x)
		case <-time.After(30 * time.Second):
			println(time.Now().Unix())
		}
	}
}
