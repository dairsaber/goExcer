package main

import (
	"fmt"
	"time"
)

//定义一个全局的channel用于通信
var channel = make(chan int)

var channel2 = make(chan int)

//定义一个打印机程序
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 1000)
	}
}

//定义两个人打印 hello 和 world
func person01() {
	printer("hello")
	channel <- 88 //向channel中写入数据 当该数据未被读取时此channel将会堵塞
}

func person02() {
	<-channel //读取channel中的数据当channel中没有写入数据此channel将会造成阻塞
	printer(" world")
	channel2 <- 89
}
func main() {
	go person01()
	go person02()
	num := <-channel2
	fmt.Printf("\n%d", num)
}
