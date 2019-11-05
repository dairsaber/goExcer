package tools

// RemoveEmptyStr 是一个去除数组中空字符串的方法！
func RemoveEmptyStr(arr []string) []string {
	out := make([]string, 0)
	for _, str := range arr {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}
