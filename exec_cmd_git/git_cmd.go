package exec_cmd_git

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//声明log为 Logger指针类型的变量
var stlog *log.Logger

// 启动cmd控制台并执行命令
func start_cmd(cmd_instruct string, cmd_args string) {
	var cmd *exec.Cmd
	// 打开cmd窗口并运行命令
	if cmd_instruct != "" {
		if cmd_args == "" {
			cmd = exec.Command(cmd_instruct)
		} else {
			cmd = exec.Command(cmd_instruct, cmd_args)
		}
	} else {
		log.Fatal("The command instruct is emputy!")
	}

	// 获取接收命令行返回值对象
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		fmt.Println("asd")
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令,这里可以区分使用
	// Start执行不会等待命令完成，Run会阻塞等待命令完成
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}

// 创建日志对象
func init_log() {
	stlog = log.New(os.Stdout, "Go: ", log.LstdFlags)
}

// 启动程序
func Start_cmd_command() {
	// stlog = init_log("./Log.log")
	init_log()
	var cmd_instruction string
	var cmd_args string
	// 打开cmd
	// cmd_instruction = "dir"
	// cmd_args = ""
	// start_cmd(cmd_instruction, cmd_args)
	// 打开cmd
	cmd_instruction = "cmd"
	cmd_args = ""
	start_cmd(cmd_instruction, cmd_args)
	// 添加
	cmd_instruction = "git add ."
	cmd_args = ""
	start_cmd(cmd_instruction, cmd_args)
	// 提交
	cmd_instruction = "git commit "
	cmd_args = "-m \"update\""
	start_cmd(cmd_instruction, cmd_args)
	// 更新
	cmd_instruction = "git push"
	cmd_args = ""
	start_cmd(cmd_instruction, cmd_args)
}
