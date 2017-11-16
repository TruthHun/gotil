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
