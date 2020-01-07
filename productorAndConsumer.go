package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 8; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {//将channel中的数据逐一打印出来知道主channel关闭
		fmt.Println("主程序消费中...",num)
		time.Sleep(time.Second)
	}
}
func main() {
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
}
