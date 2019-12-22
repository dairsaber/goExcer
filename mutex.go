package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//var mutex sync.Mutex
var rMutex sync.RWMutex

//func myPrinter(str string) {
//	mutex.Lock()
//	for _, c := range str {
//		fmt.Printf("%c", c)
//		time.Sleep(time.Millisecond * 500)
//	}
//	mutex.Unlock()
//}

//func person01() {
//	myPrinter("hello ")
//}
//func person02() {
//	myPrinter("world!")
//}
var globalNum int
//=======
//func read(out <-chan int, idx int) {
//		rMutex.RLock()
//		num := <-out
//		fmt.Printf("read ==> %d %dth\n", num, idx)
//		rMutex.RUnlock()
//}
//
//func write(in chan<- int) {
//	rMutex.Lock()
//	num := rand.Intn(1000)
//	time.Sleep(time.Second)
//	in <- num
//	fmt.Printf("write ==> %d\n", num)
//	rMutex.Unlock()
//}
//
//
func read(idx int) {
	for i := 0; i < 10; i++ {
		rMutex.RLock()
		num := globalNum
		fmt.Printf("read ==> %d %dth\n", num, idx)
		time.Sleep(time.Millisecond * 100)
		rMutex.RUnlock()
	}

}

func write() {
	rMutex.Lock()
	num := rand.Intn(1000)
	globalNum = num
	fmt.Printf("write ==> %d\n", num)
	time.Sleep(time.Millisecond * 300)
	rMutex.Unlock()
}

func main() {
	quit := make(chan bool)
	//
	//go person02()
	//go person01()
	//以上是互斥锁

	// 以下造成 隐式死锁 channel有锁定功能只要把共享数据用全局变量代替即可
	//dataChannel := make(chan int)
	//for i := 0; i < 5; i++ {
	//	go read(dataChannel, i+1)
	//}
	//for i := 0; i < 5; i++ {
	//	go write(dataChannel)
	//}

	for i := 0; i < 5; i++ {
		go read(i + 1)
	}
	for i := 0; i < 5; i++ {
		go write()
	}

	go func() {
		time.Sleep(time.Second * 10)
		<-quit
	}()
	quit <- true

}
