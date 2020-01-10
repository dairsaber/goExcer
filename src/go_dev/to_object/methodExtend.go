package main

import "fmt"
//Person 人类
type Person struct {
	name string
	age  int
}

func (p Person) sayHi() {
	fmt.Printf("Hi~ my name is %s, %d years old.\n",p.name,p.age)
}
//Student 学生类
type Student struct {
	Person
	studentNO string
}

//ExtendMethodDemoFunc 方法继承
func ExtendMethodDemoFunc() {
	aPerson := Person{name: "张三", age: 25}
	aPerson.sayHi()
	aStudent := Student{Person{name:"小俊俊",age:22}, "23252156215"}
	aStudent.sayHi()
}
