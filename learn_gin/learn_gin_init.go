package learn_gin

import (
	"github.com/gin-gonic/gin"
	"go_day/learn_gin/api"
	"go_day/learn_mysql"
	"net/http"
)

//type SqlUser struct {
//	User_mail string 'json:"user_mail" from:"user_mail"' // 问题出在使用 ·而不是‘
//	User_pass string 'json:"user_pass" from:"user_pass"'
//} // 这个结构体可以自己生成？

type SqlUser struct {
	User_mail string `json:"user_mail" from:"user_mail"`
	User_pass string `json:"user_pass" from:"user_pass"`
	//User_email interface{}
}

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
	r.POST("/add", api.Update)
	r.DELETE("/delete", api.Delete)

	err := r.Run(":0909")
	if err != nil {
		return
	}
}
