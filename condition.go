package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

//producer01 生产者函数
func producer01(in chan<- int, idx int) {
	for {
		cond.L.Lock()
		for len(in) == 5 {
			cond.Wait()
		}
		randNum := rand.Intn(2000)
		in <- randNum
		cond.L.Unlock()
		cond.Signal()
		fmt.Printf("当前生产者%d号生产的数字为：%d\n", idx, randNum)
		time.Sleep(time.Millisecond * 300)
	}
}

//customer01
func customer01(out <-chan int, idx int) {
	for {
		cond.L.Lock()
		for len(out) == 0 {
			cond.Wait()
		}
		num := <-out
		cond.L.Unlock()
		cond.Signal()
		fmt.Printf("当前消费者%d号消费的数字为：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cond.L = new(sync.Mutex)
	goodsCh := make(chan int, 5)//当这边没有给5个缓冲区的话或者和上边的cond条件不一致的话就可能出现死锁

	quit := make(chan bool)

	for i := 0; i < 5; i++ {
		go producer01(goodsCh, i)

	}
	for i := 0; i < 5; i++ {
		go customer01(goodsCh, i)
	}
	quit <- true
}
