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

func saveClient(name string, addr string) (currentClient *Client) {
	currentClient = new(Client)
	currentClient.Name = name
	currentClient.Addr = addr
	currentClient.C = make(chan string)
	clientList[addr] = *currentClient
	return
}
func MakeMsg(clnt *Client, content string) (buf string) {
	buf = "[" + clnt.Name + "]" + content+"\n"
	return

}

func WriteToClient(client *Client, conn net.Conn) {
	for {
		msg := <-client.C
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发送给客户端出错！")
			continue
		}
	}
}

func HandlerClientConnect(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	currentClient := saveClient(remoteAddr, remoteAddr)
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
			RemoveClient(remoteAddr)
		case <-hasData: //重置下面的定时器
		case <-time.After(time.Second * 60):
			RemoveClient(remoteAddr)
			message <- MakeMsg(currentClient, "长时间为发言已被提出群聊")
			return
		}
	}

}
func ReName(client *Client, newName string) (oldName string) {
	oldName = client.Name
	client.Name = newName
	return
}

func RemoveClient(key string) {
	delete(clientList, key)
}

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
