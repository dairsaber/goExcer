package main

import (
	"fmt"
)

//PI 圆周率
const PI float32 = 3.141592654

//Rectangle 定义一个长方形的结构体类型]
type Rectangle struct {
	width, height float32
}

//Circle 定义选型结构体类型
type Circle struct {
	r float32
}

func (r Rectangle) area() float32 {
	return r.height * r.width
}

func (c Circle) area() float32 {
	return PI * (c.r) * (c.r)
}
func main() {
	rectangle01 := Rectangle{width: 11.2, height: 32.52}
	fmt.Println("rectangle01==>width:11.2,height:32.52,area is ===>", rectangle01.area())
	circle01 := Circle{r: 25.63}
	fmt.Println("circle01==>r:25.63,area is ===>", circle01.area())
	//======复杂盒子面向对象的案例======
	ComplexObjectFunc()
	ExtendMethodDemoFunc()
}
