package learn

import (
	"fmt"
	"time"
)

func Day2() {
	ch1 := make(chan int)
	go fmt.Println(<-ch1)
	ch1 <- 5
	time.Sleep(1 * time.Second)
}
