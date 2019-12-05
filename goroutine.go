package main

import (
	"fmt"
	"runtime"
	"time"
)

func gotest01() {
	for i := 0; i < 5; i++ {
		fmt.Println("i am gotest01")
		time.Sleep(time.Millisecond * 100)
	}
}

func gotest02() {
	for i := 0; i < 5; i++ {
		fmt.Println("i am test02")
		time.Sleep(time.Second)
	}
}
func main() {

	go gotest01()
	go gotest02()

	for i := 0; i < 10; i++ {
		runtime.Gosched() //出让时间轮片
		fmt.Println("Hi iam a main progress")
		time.Sleep(time.Second)
	}
}
