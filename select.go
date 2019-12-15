package main

import (
	"fmt"
	"runtime"
	"time"
)

/*实现一个斐波那契数列*/

//用来监听输出数字
func printResult(out <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-out:
			fmt.Printf("%d ", num)
		case <-quit:
			runtime.Goexit()
		}
	}
}

//测试两个变量交换
func change() {
	a, b := "a", "b"
	fmt.Println("a===>", a, "b===>", b)
	a, b = b, a
	fmt.Println("a===>", a, "b===>", b)

}
func main() {
	change() // 事实证明go在交换两个变量的值时是不需要声明第三个变量的
	dataCh := make(chan int)
	quit := make(chan bool)
	go printResult(dataCh, quit)
	x, y := 1, 1
	println("斐波那契数列===>")//go 的io 打印是异步的 所以不一定先打这行字让其休眠一段时间
	time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		x, y = y, x+y
		dataCh <- x
		time.Sleep(time.Second)
	}
	quit <- true
}
