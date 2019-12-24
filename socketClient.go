package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ioBuffer := make([]byte, 4096)
	fmt.Print("请输入连接配置\n>")
	n, err := os.Stdin.Read(ioBuffer)
	if err != nil {
		fmt.Println("system error")
		return
	}
	config := strings.Trim(string(ioBuffer[:n]),"\r\n")
	conn, err := net.Dial("tcp", config)
	if err != nil {
		fmt.Println("client config error")
		return
	}
	defer conn.Close()

	go func() {
		writeBuff := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(writeBuff)
			if err != nil {
				fmt.Println("system error")
				continue
			}
			//str := string(writeBuff[:n])
			conn.Write(writeBuff[:n])
		}
	}()
	readBuff := make([]byte, 4096)
	for {
		n, err := conn.Read(readBuff)
		if n == 0 {
			return
		}
		if err != nil {
			continue
		}
		fmt.Println(string(readBuff[:n]))
	}
}
