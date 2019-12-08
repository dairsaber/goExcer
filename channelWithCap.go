package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子go程", len(ch), cap(ch))//打印是IO操作，所以在go程中打印混乱是正常的
		}
	}()
	time.Sleep(time.Second * 3)
	for i := 0; i < 8; i++ {
		num := <-ch
		fmt.Println("主go程", num)

	}
}
