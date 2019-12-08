package main

import (
	"fmt"
	"time"
)

//sendCh 单向写入channel
func sendCh(ch chan<- int) {
	close(ch)
}

//receiveCh 单向读取channel
func receiveCh(ch <-chan int) int {
	num := <-ch
	fmt.Println(num)
	return num
}

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		sendCh(ch)
	}()

	for {
		//当能够读到数据时终止主go程
		num := receiveCh(ch)
		if num == 0 {
			fmt.Println("程序执行完毕！")
			break
		}
	}

}
