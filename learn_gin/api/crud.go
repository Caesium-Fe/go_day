package api

import (
	"github.com/gin-gonic/gin"
	"go_day/learn_mysql"
	"log"
	"net/http"
	"strconv"
)

type SqlUser struct {
	//id int `json:"id" from:"id"`
	User_mail string `json:"user_mail" from:"user_mail"`
	User_pass string `json:"user_pass" from:"user_pass"`
	//User_email interface{}
}

// Save
/*
http://127.0.0.1:909/save
{
    "user_mail": "123@11.com",
    "user_pass": "123"
}
*/
func Save(c *gin.Context) {
	var user SqlUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := learn_mysql.Db.Exec("insert into user(email, password) values(?,?)", user.User_mail, user.User_pass)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "save success!"})
}

// SelectById
/*
http://127.0.0.1:909/select?id=1
*/
func SelectById(c *gin.Context) {
	var user SqlUser
	id, _ := strconv.Atoi(c.Query("id"))
	//id = int(id)
	err := learn_mysql.Db.QueryRow("select email, password from user where id=?", id).Scan(&user.User_mail, &user.User_pass)
	//err := learn_mysql.Db.QueryRow("select id, email, password from user where id=?", id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"user-mail": nil,
			"user-pass": nil,
		})
		log.Fatalln(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user-mail": user.User_mail,
		"user-pass": user.User_pass,
	})
}

// Update
/*
http://127.0.0.1:909/update
{
   "id":"1",
   "email":"cheng1hao",
   "password":19
}
*/
func Update(c *gin.Context) {
	var user SqlUser
	id, _ := strconv.Atoi(c.Query("id"))
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := learn_mysql.Db.Exec("update user set email=?, password=? where id=?", user.User_mail, user.User_pass, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "Update success!"})
}

// Delete
/*
http://127.0.0.1:909/delete?id=1
*/
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	result, err := learn_mysql.Db.Exec("delete from user where id=?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "Delete success!"})
}
