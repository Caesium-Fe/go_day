package learn

import (
	"fmt"
	"log"
	"os"
)

//使用接口的方式实现一个既可以往终端写日志也
//可以往文件写日志的简易日志库。

// 写入终端结构体
type Ultimate struct {
	Prefix string
	logln  string
}

// 写入文件结构体
type Files struct {
	Prefix string
	logln  string
}

// 写日志接口
type LogWrite interface {
	writeLog()
}

// 实现终端接口
func (u *Ultimate) writeLog() {
	logger := log.New(os.Stdout, u.Prefix, log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println(u.logln)
}

// 实现文件接口
func (f *Files) writeLog() {
	logFile, err := os.OpenFile("./zbc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println(f.logln)
	log.SetPrefix(f.Prefix)
}

func Use_interface_to_log() {
	// var u Ultimate
	// var f Files
	var l LogWrite
	l = &Ultimate{
		Prefix: "[我嫩爹]",
		logln:  "ding nei ge fai",
	}
	l.writeLog()
	l = &Files{Prefix: "[我嫩爹]",
		logln: "ding nei ge fai"}
	l.writeLog()

}

/*
// 相对完整的一个练习解答
// 比我好在顾及的更广

// 接口
type Logger interface {
	consoleLog()  // 终端
	fileLog()     // 文件
}

// 用户结构体
type User struct {
	username string
	password string
}

// User实现方法
func (u User) consoleLog() {
	t := time.Now()
	fmt.Printf("用户创建成功！用户名为：%s", u.username)
	fmt.Printf("用户密码是：%s\n", u.password)
	fmt.Printf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func (u User) fileLog() {
	t := time.Now()
	file, err := os.OpenFile("./"+u.username+".txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", u.username) + "密码是：" + u.password + "\n" + fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	_, _ = file.WriteString(data)
	_ = file.Close()
}

// User中的字段初始化
func newUser(username, password string) User {
	return User{
		username: username,
		password: password,
	}
}

// 创建用户对象
func createUser() {
	var (
		username string
		password string
	)
	fmt.Print("请输入用户名：")
	_, err := fmt.Scan(&username)
	fmt.Print("请输入一个密码：")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println("输入错误！！ERROR:", err)
	}
	u := newUser(username, password)
	u.consoleLog()
	u.fileLog()
}

// 主函数
func main() {
	createUser()
}

*/
