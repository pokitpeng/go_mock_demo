package utils

// HasString 检查切片是否存在某个字符串
func HasString(array []string, elem string) bool {
	for _, v := range array {
		if v == elem {
			return true
		}
	}
	return false
}
