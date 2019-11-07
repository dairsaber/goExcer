package main

import (
	"fmt"
	"reflect"

	"./base"
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

	waitRemoveArr := []string{"ckjcf", "sdads", "remove", "cssfre", "end"}
	removedData := tools.RemoveByIndex(waitRemoveArr, 2)
	fmt.Println("移除完之后的值===>", removedData, waitRemoveArr)

	testStr := "wo shi yi zhi lai zi bei fang fang fang de lang , ta ye shi yi zhi lai zi nan fang de gou!"
	resultMap := base.CountWords(testStr)
	fmt.Println("这是map的一个练习===>", resultMap)

	//创建一个test.txt的文件
	const fileName string = "/Users/panafeng/project/demo/Go/demo/test.txt"
	success := base.CreateFile(fileName)
	if success {
		myMap := map[string]string{"name": "heheda", "hobby": "喜欢吃饭睡觉打豆"}
		readSuccess := base.WriteFile(myMap, fileName)
		if readSuccess {
			person := base.CreateStructFromFile(fileName)
			fmt.Println("person===>", *person)
		}
	}
}
