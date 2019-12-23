package main

import (
	"fmt"
	"net"
)

func main() {
	//配置资源socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listener 建立失败==>", err)
		return
	}
	fmt.Println("服务器等待连接中...")
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept 建立连接失败===>", err)
		return
	}
	fmt.Println("已经和客户端建立了连接，等待客户端发送数据！")

	defer conn.Close()
	buff := make([]byte, 4096)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("读取客户端数据失败！===>", err)
	}
	fmt.Println("客户端数据为：", string(buff[:n]))
}
