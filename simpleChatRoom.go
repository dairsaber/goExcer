package main

import (
	"fmt"
	"net"
	"time"
)

var message = make(chan string)

//Client 客户端对象结构体
type Client struct {
	Name string
	Addr string
	C    chan string
}

var clientList map[string]Client
//SaveClient 保存对应的客户端
func SaveClient(name string, addr string) (currentClient *Client) {
	currentClient = new(Client)
	currentClient.Name = name
	currentClient.Addr = addr
	currentClient.C = make(chan string)
	clientList[addr] = *currentClient
	return
}
//MakeMsg 生成对应的消息
func MakeMsg(clnt *Client, content string) (buf string) {
	buf = "[" + clnt.Name + "]" + content+"\n"
	return

}
// WriteToClient 将详细写给客户端
func WriteToClient(client *Client, conn net.Conn) {
	// for {
	// 	msg := <-client.C
	// 	_, err := conn.Write([]byte(msg))
	// 	if err != nil {
	// 		fmt.Println("发送给客户端出错！")
	// 		continue
	// 	}
	// }
	for msg := range client.C {//简便写法
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发送给客户端出错！")
			return
		}
	}
}
//HandlerClientConnect 客户端连接的处理函数
func HandlerClientConnect(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	currentClient := SaveClient(remoteAddr, remoteAddr)
	currentMsg := MakeMsg(currentClient, "login")
	isQuit := make(chan bool)
	hasData := make(chan bool)
	go WriteClientMsg(conn, currentClient, isQuit, hasData)
	go WriteToClient(currentClient, conn)
	message <- currentMsg

	//保证此go程在未退出时不要结束则采用select监听退出
	for {
		select {
		case <-isQuit:
			RemoveClient(remoteAddr,currentClient)
		case <-hasData: //重置下面的定时器
		case <-time.After(time.Second * 60):
			message <- MakeMsg(currentClient, "长时间为发言已被提出群聊")
			RemoveClient(remoteAddr,currentClient)
			return
		}
	}

}
//ReName 更新客户端昵称.
func ReName(client *Client, newName string) (oldName string) {
	oldName = client.Name
	client.Name = newName
	return
}
//RemoveClient 移除客户端
func RemoveClient(key string,client *Client) {
	//关闭通道这样会让使用改通道的go程终止并推出
	close(client.C)
	delete(clientList, key)
}
//WriteClientMsg 生成对应操作的message，并将message存入全局变量中
func WriteClientMsg(conn net.Conn, currentClient *Client, isLogout chan<- bool, hasData chan<- bool) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			isLogout <- true
			message <- MakeMsg(currentClient, "is Logout!")
			return
		}
		if err != nil {
			fmt.Println("服务器出错！")
			continue
		}
		currentMsg := string(buf[:n-1])
		if len(currentMsg) > 6 && string(currentMsg[:7]) == "rename " {
			currentName := string(currentMsg[7:])
			oldName := ReName(currentClient, currentName)
			message <- MakeMsg(currentClient, oldName+"更新昵称成功！")
		} else {
			message <- MakeMsg(currentClient, currentMsg)
		}
		hasData <- true
	}
}
//Manager 监视全局变量message channel 并将消息取出发送给在线的客户端
func Manager() {
	clientList = make(map[string]Client)
	for {
		msg := <-message
		for _, client := range clientList {
			client.C <- msg
		}
	}
}

func main() {
	//新建一个socket套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("socket 创建失败！！")
	}
	go Manager()
	fmt.Println("服务端启动成功！在127.0.0.1:8001端口上")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("socket 创建失败！！")
		}
		go HandlerClientConnect(conn)

	}

}
