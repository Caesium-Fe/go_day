package learn

func Deferl() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok { //判断通道是否关闭
				return
			}
			println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done

	donr := make(chan struct{})
	cr := make(chan int)
	go func() {
		defer close(donr)
		for x := range cr { //循环获取消息
			println(x)
		}

	}()
	cr <- 1
	cr <- 2
	cr <- 3
	close(cr)
	<-donr
}

//及时使用close函数关闭通道引发结束通知，避免出现可能的死锁
