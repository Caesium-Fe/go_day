package learn

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Print系列函数会将内容输出到系统的标准输出，
// 区别在于Print函数直接输出内容，
// Printf函数支持格式化输出字符串，
// Println函数会在输出内容的结尾添加一个换行符。
func Easy_test() {
	fmt.Print("在终端打印该信息。")
	name := "沙河小王子"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

// Fprint系列函数会将内容输出到
// 一个io.Writer接口类型的变量w中，
// 我们通常用这个函数往文件中写入内容。
func Fprint_test() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./learn_Fprint.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

// Sprint系列函数会把传入的数据生成
// 并返回一个字符串。
func Sprint_test() {
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)
}

//Errorf函数根据format参数生成格式化字符串
// 并返回一个包含该字符串的错误。
func Errorf_test() {
	//通常使用这种方式来自定义错误类型，例如：
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)
	//Go1.13版本为fmt.Errorf函数新加了一个%w占位符用
	// 来生成一个可以包裹Error的Wrapping Error。
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	if w != nil {
		fmt.Println(w)
	}
}

func Scan_test() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println("please use key input: name(string),age(int),married(bool)")
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

//fmt.Scanf不同于fmt.Scan简单的
//以空格作为输入数据的分隔符，
// fmt.Scanf为输入数据指定了具体的输入内容格式，
//只有按照格式输入数据才会被扫描并存入对应变量。
func Scanf_test() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// fmt.Scanln遇到回车就结束扫描了，这个比较常用。
func Scanln_test() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// 有时候我们想完整获取输入的内容，
// 而输入的内容可能包含空格，
// 这种情况下可以使用bufio包来实现。
// 示例代码如下：
func BufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
