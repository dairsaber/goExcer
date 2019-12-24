package main

import (
	"fmt"
	"net"
	"strings"
)

func handlerConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端已连接")
	buff := make([]byte, 4096)
	for {
		n, err := conn.Read(buff)
		if n == 0 {
			fmt.Println(addr, "客户端已经关闭")
			return
		}
		if err != nil {
			fmt.Println("读取数据失败！")
			return
		}
		content := string(buff[:n])
		fmt.Println(content)
		if strings.Trim(content, "\n") == "exit" {
			fmt.Println(addr, "客户端请求关闭！")
			conn.Write([]byte("退出连接"))
			return
		}
		conn.Write([]byte(strings.ToUpper(content) ))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("服务器创建不成功！")
		return
	}
	fmt.Println(listener.Addr(), "服务端已启动")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败！")
		}
		go handlerConn(conn)
	}
}
