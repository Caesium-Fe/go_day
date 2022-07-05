package learn

import (
	"fmt"
	"sync"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)
func Mutex1() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	//fmt.Println(counter)
	wg.Wait()
	fmt.Println(counter)
}
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			value++
			counter = value
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}