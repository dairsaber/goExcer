package main

import (
	"fmt"
	"reflect"

	"./tools"
)

func main() {
	fmt.Println("demo")
	tools.MyPrint("这是tools包中的方法")

	testStrSlice := []string{"heheda", "", "", "dadahe", "wolegevao"}
	result := tools.RemoveEmptyStr(testStrSlice)
	fmt.Println("去除数组中的空字符串", result)

	testStrSlice2 := []string{"hehehe", "", "", "dda", "red", "yellow", "banana", "red", "dda"}
	result2 := tools.QuChong(testStrSlice2)
	fmt.Println("去重结果===>", result2)

	fmt.Println("=====", reflect.TypeOf(""))
}
