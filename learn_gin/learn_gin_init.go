package learn_gin

import (
	"github.com/gin-gonic/gin"
	"go_day/learn_gin/api"
	"go_day/learn_mysql"
	"net/http"
)

func GinInit() {
	r := gin.Default()
	//var u SqlUser
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	learn_mysql.LinkSql()

	r.PUT("/save", api.Save)
	r.GET("/select", api.SelectById)
	r.GET("/showAll", api.ShowAll)
	r.POST("/add", api.Update)
	r.DELETE("/delete", api.Delete)

	r.GET("/redirectTest", api.RedirectTest)
	r.GET("/redirectURL", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/redirectURL2"
		r.HandleContext(c)
	})
	r.GET("/redirectURL2", api.RedirectURL2)

	r.GET("/goroutine1", api.Goroutine1)
	r.GET("/goroutine2", api.Goroutine2)

	err := r.Run(":9092")
	if err != nil {
		return
	}
}
