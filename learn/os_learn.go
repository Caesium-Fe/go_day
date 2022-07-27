package learn

import (
	"fmt"
	"io"
	"os"
)

func OpenAndClose() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("E:/go_day/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	fmt.Println(file)
	// 关闭文件
	file.Close()
}

func Read_file() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("E:/go_day/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	// var tmp []byte
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

func For_read_file() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("E:/go_day/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	// 循环读取文件时，虽然一次也只能读取128字节
	// 但每次循环后都被添加到content里
	// 这样就可以将文件所有内容都读取到
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
