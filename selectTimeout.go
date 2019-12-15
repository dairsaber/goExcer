package main

import (
	"fmt"
	"time"
)

func getData(ch chan<- int) {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Duration(i)*time.Second)//这边控制一下时间
		ch <- i
	}
	close(ch)
}

func listenner(dataCh <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-dataCh:
			if num == 0 {
				fmt.Println("finished")
				<-quit
				break
			}
			fmt.Println("current data is ===>", num)
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			<-quit
			return
		}
	}
}

func main() {
	datach := make(chan int)
	quit := make(chan bool)
	go listenner(datach, quit)
	go getData(datach)
	quit <- true
}
