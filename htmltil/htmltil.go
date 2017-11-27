package htmltil

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

//html压缩【非gzip压缩，只是去掉多余的空格或者是换行和tab】
//@param			HtmlStr				需要压缩的html字符串
//@param			DoesTrimAllSpace	是否存在预览的代码，如果html页面中存在code，如python代码等，则不能替换掉空格，所以这里要慎重
//@return			NewHtmlStr			处理后的html字符串
func Compress(HtmlStr string, DoesTrimAllSpace ...bool) (NewHtmlStr string) {
	if len(DoesTrimAllSpace) > 0 && DoesTrimAllSpace[0] {
		r, _ := regexp.Compile(">\\s{1,}<")
		HtmlStr = r.ReplaceAllString(HtmlStr, "><")
	}
	HtmlStr = strings.Replace(HtmlStr, "\n", "", -1)
	HtmlStr = strings.Replace(HtmlStr, "\t", "", -1)
	HtmlStr = strings.Replace(HtmlStr, "\r", "", -1)
	HtmlStr = strings.Replace(HtmlStr, "\n\r", "", -1)
	return HtmlStr
}

//golang调用浏览器打开指定链接
//@param			uri				需要打开的url链接地址
//@return			err				错误
func OpenByBrowser(uri string) (err error) {
	var cmds = map[string]string{"windows": "cmd /c start", "darwin": "open", "linux": "xdg-open"}
	if run, ok := cmds[runtime.GOOS]; !ok {
		return errors.New(fmt.Sprintf("找不到当前平台(%v)调用浏览器打开链接的指令", runtime.GOOS))
	} else {
		return exec.Command(run, uri).Start()
	}
}

//url的query请求解析，如：http://example.com?name=truthhun&age=18&hobbies[]=football&hobbies[]=swimming，则这里的QueryStr为问号后面的部分
//@param		QueryStr		请求字符串
//@return		params			解析后的参数，如果有select或者checkbox等多选的情况下，则interface{}为[]string类型，否则都是string，在使用的时候，用类型断言即可
//需要注意的是，如果请求参数中的参数名带有“[]”，那么解析后的参数中也是带有中括号的
//功能类似url.ParseQuery
func ParseUrlQuery(QueryStr string) (params map[string]interface{}) {
	var kvs = make(map[string][]string)
	slice := strings.Split(QueryStr, "&")
	params = make(map[string]interface{})
	if len(slice) > 0 {
		for _, q := range slice {
			if param := strings.Split(q, "="); len(param) == 2 {
				if strings.HasSuffix(param[0], "[]") { //单个key存在多值
					kvs[param[0]] = append(kvs[param[0]], param[1])
				} else {
					params[param[0]] = param[1]
				}
			}
		}
	}
	if len(kvs) > 0 {
		for k, v := range kvs {
			params[k] = v
		}
	}
	return
}
