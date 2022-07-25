package git_add

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func start_cmd() {
	cmd := exec.Command("cmd")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	go func() {
		//写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("git add .\ngit commit -m \"update\"\ngit push\n")
	}()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	log.Println("1")
	fmt.Println(out.String())

	// cmd := exec.Command("cmd", "/C", "dir")
	// cmd := exec.Command("cmd", "/C", "git add .")
	// cmd.Dir = "E:/go_day"
	// fmt.Println(cmd.Dir)
	// stdout, err := cmd.StdoutPipe()
	// // fmt.Println(err)
	// if err != nil { //获取输出对象，可以从该对象中读取输出结果
	// 	// fmt.Println("1")
	// 	log.Fatal(err)
	// }
	// defer stdout.Close()                // 保证关闭输出流
	// if err := cmd.Start(); err != nil { // 运行命令
	// 	log.Fatal(err)
	// 	// fmt.Println("2")
	// }
	// if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
	// 	log.Fatal(err)
	// 	// fmt.Println("3")
	// } else {
	// 	log.Println(string(opBytes))
	// 	// fmt.Println("4")
	// }

	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("error:", err)
	// }
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(output))
}

// 获取当前运行的路径
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取当前执行程序所在的绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func Src() {
	a := getCurrentDirectory()
	fmt.Println(a)
	start_cmd()
}
