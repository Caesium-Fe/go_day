package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Goroutine1(c *gin.Context) {
	// 拷贝一份副本在Goroutine中使用
	tmp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		// 这里使用的是值拷贝的tmp
		log.Println("routine1 finish, url path is:" + tmp.Request.URL.Path)
	}()
}

func Goroutine2(c *gin.Context) {
	time.Sleep(5 * time.Second)
	// 因为没有使用 goroutine， 不需要拷贝上下文
	log.Println("routine1 finish, url path is:" + c.Request.URL.Path)
}
