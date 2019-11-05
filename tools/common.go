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
