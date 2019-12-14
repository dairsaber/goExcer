package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	//NewTicker定时器的一般用法
	myTimer := time.NewTimer(time.Second * 3)
	ch := myTimer.C
	currentTime := time.Now().Second()
	fmt.Println(currentTime)
	<-ch //读取定时器channel中的值抛弃，让程序解除阻塞继续向下走
	fmt.Println("定时器完成", time.Now().Second())

	//After 相当于上面的简化版直接返回C
	afterTimer := time.After(time.Second * 3)
	<-afterTimer
	fmt.Println(time.Now().Second())

	//中断定时器
	myTimer2 := time.NewTimer(time.Second * 5)
	go func() {
		<-myTimer2.C
		fmt.Println("我是子go程！")
	}()
	time.Sleep(time.Second * 2)
	myTimer2.Stop()
	fmt.Printf("我是主go程中断myTimer2，所以子go程中无法执行，也就是不会打印那段《我是子go程！》的文字！")
	//循环定时
	times := 0
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			nowTime := <-ticker.C
			fmt.Println("nowTime is ", nowTime.Second(), times)
			if times == 8 {
				ticker.Stop()
				quit <- false //然主程序一直阻塞 ,不销毁子go程
			}
			times++
		}
	}()
	<-quit
}
