package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}
	defer conn.Close()
	conn.Write([]byte("Are you ready?"))

	buff := make([]byte, 4096)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("读取服务端数据失败")
		return
	}
	fmt.Println("服务器返回：", string(buff[:n]))
}
