package tools

//QuChong 去重
func QuChong(arr []string) []string {
	out := make([]string, 0)
	for _, str := range arr {
		if !includes(out, str) {
			out = append(out, str)
		}
	}
	return out
}

// includes 判断数组中是否存在某项
func includes(arr []string, item string) bool {
	for _, str := range arr {
		if str == item {
			return true
		}
	}
	return false
}

//RemoveByIndex 移除数组中的某个项
func RemoveByIndex(arr []string, index int) []string {
	//这种写法原切片会受到影响
	// before := arr[:index]
	// after := arr[1:]
	// copy(after, before)
	// return after

	//下面这种写法也会会改变原切片
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}
