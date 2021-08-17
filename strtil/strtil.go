package strtil

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/axgle/mahonia"
	"github.com/gogs/chardet"
	"github.com/rogpeppe/go-charset/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

//首字母大写
//@param			str			需要对首字母进行大写的字符串
//@return						范湖
func UpperFirst(str string) string {
	return strings.Replace(str, str[0:1], strings.ToUpper(str[0:1]), 1)
}

// ConvertToUTF8 将任意可能非utf8字符串转为utf8编码
func ConvertToUTF8(char string) (newChar string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			return
		}
	}()

	newChar = strings.TrimSpace(char)
	if newChar == "" {
		return
	}

	if utf8.ValidString(newChar) {
		return
	}

	// 检测字符编码
	enc := detectCharacter([]byte(newChar))
	newChar, err = encodeUTF8(newChar, enc)
	if err != nil {
		return
	}

	newChar = mahonia.NewDecoder(strings.ToLower(enc)).ConvertString(newChar)
	return
}

func detectCharacter(cont []byte) string {
	var res *chardet.Result
	res, err := chardet.NewTextDetector().DetectBest(cont)
	if err != nil {
		return "gbk"
	}
	return res.Charset
}

//convert GBK to UTF-8
func encodeUTF8(cont, enc string) (str string, err error) {
	var b []byte
	in := strings.NewReader(cont)
	switch strings.ToUpper(enc) {
	case "GB18030", "GB-18030":
		b, err = ioutil.ReadAll(transform.NewReader(in, simplifiedchinese.GB18030.NewDecoder()))
	case "GBK", "GB2312", "GB-2312":
		b, err = ioutil.ReadAll(transform.NewReader(in, simplifiedchinese.GBK.NewDecoder()))
	case "HZGB2312", "HZ-GB2312":
		b, err = ioutil.ReadAll(transform.NewReader(in, simplifiedchinese.HZGB2312.NewDecoder()))
	case "BIG":
		b, err = ioutil.ReadAll(transform.NewReader(in, traditionalchinese.Big5.NewDecoder()))
	case "UTF8", "UTF-8":
		str = cont
		return
	default:
		var rio io.Reader
		rio, err = charset.NewReader(enc, strings.NewReader(cont))
		if err != nil {
			return
		}
		b, err = ioutil.ReadAll(rio)
	}
	str = string(b)
	return
}
