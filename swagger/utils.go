package swagger

//IsInStringList 判断字符串是否在stirng slice 里
func IsInStringList(list []string, s string) bool {
	for i := range list {
		if list[i] == s {
			return true
		}
	}
	return false
}
