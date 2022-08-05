package learn

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)
	for {
		// reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}

	// 黏包服务端core
	// defer conn.Close()
	// reader := bufio.NewReader(conn)
	// var buf [1024]byte
	// for {
	// 	n, err := reader.Read(buf[:])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read from client failed, err:", err)
	// 		break
	// 	}
	// 	recvStr := string(buf[:n])
	// 	fmt.Println("收到client发来的数据：", recvStr)
	// }

}

func Tcp_server() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}

	// 黏包服务端core
	// listen, err := net.Listen("tcp", "127.0.0.1:30000")
	// if err != nil {
	// 	fmt.Println("listen failed, err:", err)
	// 	return
	// }
	// defer listen.Close()
	// for {
	// 	conn, err := listen.Accept()
	// 	if err != nil {
	// 		fmt.Println("accept failed, err:", err)
	// 		continue
	// 	}
	// 	go process(conn)
	// }
}
