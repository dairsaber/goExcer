package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	//创建serve地址资源
	serverAddr,err :=  net.ResolveUDPAddr("udp","127.0.0.1:8888")
	if err !=nil {
		fmt.Println("ResolveUDPAddr ==> error",err)
	}
	udpConn,err := net.ListenUDP("udp",serverAddr)
	if err !=nil {
		fmt.Println("创建udp出错",err)
	}
	defer udpConn.Close()
 fmt.Println("udp服务器启动成功",serverAddr)

	buff := make([]byte,4096)
	n,clientAddr,err :=  udpConn.ReadFromUDP(buff)
	if err!=nil {
		fmt.Println("读取客户端数据出错",err)
	}
	fmt.Printf("接收到%v发送过来的%s数据",clientAddr,string(buff[:n]))

	//回写给客户端

	message := time.Now().String()
	_, err = udpConn.WriteToUDP([]byte(message), clientAddr)
	if err != nil {
		fmt.Println("服务端回写数据失败==>",err)
	}

}
//udp 的客户端和tcp差不多只是协议写成udp就可以了
