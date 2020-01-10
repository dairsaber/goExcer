package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	//WHITE is a color 白色
	WHITE = iota
	//RED 红色
	RED
	//BLACK 黑色
	BLACK
	//GREEN 绿色
	GREEN
	//BLUE 蓝色
	BLUE
)

//Color 颜色
type Color byte

// Box 盒子结构体
type Box struct {
	width, height, length float64
	color                 Color
}
//Volumn 获得盒子的体积
func (b *Box) Volumn() float64 {
	return b.height * b.length * b.width
}

//BoxList 存储盒子列表
type BoxList []*Box

//GetMaxVolBox 获得体积最大的 box 引用
func (bl BoxList) GetMaxVolBox() *Box {
	v := 0.0
	var tempBox *Box
	for _, box := range bl {
		if currentVol := box.Volumn(); currentVol > v {
			tempBox = box
			v = currentVol
		}
	}
	return tempBox
}

//SetColor 给盒子设置颜色
func (b *Box) SetColor(c Color) {
	b.color = c
}
//GetColorName 获取盒子的颜色的名称
func(b *Box) GetColorName() string{
colorNames := []string{"WHITE","RED",	"BLACK","GREEN","BLUE"}
	return colorNames[byte(b.color)]
}
//CreateBoxes 随机生成几个盒子
func CreateBoxes() BoxList {
	boxLength := 3 + rand.Intn(7)
	bl := make([]*Box, 0)
	for i := 0; i < boxLength; i++ {
		width := 5.0 + rand.Float64()*10
		height := 5.0 + rand.Float64()*10
		length := 5.0 + rand.Float64()*10
		box := new(Box)
		box.height = height
		box.width = width
		box.length = length
		bl = append(bl, box)
	}
	return BoxList(bl)
}

//ComplexObjectFunc 复杂的对象练习
func ComplexObjectFunc() {
	rand.Seed(time.Now().UnixNano())
	boxes := CreateBoxes()
	volumMaxBox := boxes.GetMaxVolBox()
	fmt.Println("volumMaxBox===>", volumMaxBox)
	fmt.Println("随机盒子===>", boxes)
	fmt.Println("最大盒子的体积", volumMaxBox.Volumn())
	fmt.Println("最大盒子的颜色", volumMaxBox.GetColorName())
	fmt.Println("设置最大盒子的颜色为红色RED")
	volumMaxBox.SetColor(Color(RED))
	fmt.Println("最大盒子的颜色===>", volumMaxBox.GetColorName())

}
