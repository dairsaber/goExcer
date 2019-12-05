package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer fmt.Println("heheda i am test")
	runtime.Goexit() // 退出当前go程
	// return  ===> 退出当前函数

}

func main() {
	go func() {
		defer fmt.Println("i am a 匿名函数")
		test()
		for i := 0; i < 8; i++ {
			fmt.Println("i am test")
			time.Sleep(time.Second)
		}

	}()

	for {//这种写法让程序在此死循环
		time.Sleep(time.Second)
	}
}
