package strtil

import "strings"

//首字母大写
//@param			str			需要对首字母进行大写的字符串
//@return						范湖
func UpperFirst(str string) string {
	return strings.Replace(str, str[0:1], strings.ToUpper(str[0:1]), 1)
}
