package base

import "strings"

//CountWords 计算一句话中重复单词的个数
func CountWords(str string) map[string]int {
	strs := strings.Fields(str)
	temp := make(map[string]int)
	for _, str := range strs {
		if value, ok := temp[str]; ok {
			temp[str] = value + 1
		} else {
			temp[str] = 1
		}
	}
	return temp
}
