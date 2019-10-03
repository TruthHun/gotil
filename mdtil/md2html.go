package mdtil

import "github.com/russross/blackfriday"

//将markdown内容转成html内容
//@param            MarkdownContent     markdown文本内容
//@return           html                转化后的html
func Md2html(MarkdownContent string) (html string) {
	out := blackfriday.Run([]byte(MarkdownContent))
	return string(out)
}
