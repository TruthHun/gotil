package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/httplib"
)

//构造request请求。[如果需要gzip压缩，请自行在headers参数配置，但是处理响应体的时候，记得使用gzip解压缩]
//@param			method			请求方法：get、post、put、delete、head
//@param			url				请求链接
//@param			referrer		如果有referrer，则配置该选项的header.注意：这里可以置空，然后在headers参数中配置
//@param			cookie			如果有cookie，则配置该选项的header.注意：这里可以置空，然后在headers参数中配置
//@param			os				操作系统，用于配置UA，参数值：windows、linux、Android、ios、mac。默认mac下的谷歌浏览器UA
//@param			iscn			是否是中文请求，用于访问多语言的站点
//@param			isjson			是否请求的是json数据
//@param			headers			更多请求头配置项
//@return							返回http请求
func BuildRequest(method, url, referrer, cookie, os string, iscn, isjson bool, headers ...map[string]string) *httplib.BeegoHTTPRequest {
	var req *httplib.BeegoHTTPRequest
	switch strings.ToLower(method) {
	case "get":
		req = httplib.Get(url)
	case "post":
		req = httplib.Post(url)
	case "put":
		req = httplib.Put(url)
	case "delete":
		req = httplib.Delete(url)
	case "head":
		req = httplib.Head(url)
	default:
		req = httplib.Get(url)
	}

	//设置referrer
	if len(referrer) > 0 {
		req.Header("Referrer", referrer)
	}
	//设置cookie
	if len(cookie) > 0 {
		req.Header("Cookie", cookie)
	}
	//设置host
	host_slice := strings.Split(url, "://")
	if len(host_slice) > 1 {
		host := strings.Split(host_slice[1], "/")[0]
		req.SetHost(host)
	}
	//压缩
	//req.Header("Accept-Encoding", "gzip, deflate, br")
	//中文
	if iscn {
		req.Header("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	} else {
		req.Header("Accept-Language", "en-US,en;q=0.8,zh;q=0.6")
	}
	//是否是json采集
	if isjson {
		req.Header("Accept", "application/json")
		req.Header("X-Request", "JSON")
		req.Header("X-Requested-With", "XMLHttpRequest")
	} else {
		req.Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	}

	//系统设置
	switch strings.ToLower(os) {
	case "windows":
		req.Header("User-Agent", "Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36")
	case "linux":
		req.Header("User-Agent", "Mozilla/5.0 (X11; U; Linux i686) AppleWebKit/534.15 (KHTML, like Gecko) Ubuntu/10.10 Chromium/10.0.613.0 Chrome/10.0.613.0 Safari/534.15")
	case "mac":
		req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3107.4 Safari/537.36")
	case "android":
		req.Header("User-Agent", "MQQBrowser/26 Mozilla/5.0 (Linux; U; Android 2.3.7; MB200 Build/GRJ22; CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")
	case "ios":
		req.Header("User-Agent", "Mozilla/5.0(iPhone; CPU iPhone OS 9_3_3 like Mac OS X)AppleWebkit/601.1.46(KHTML,like Gecko)Mobile/13G3")
	default: //mac
		req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3107.4 Safari/537.36")
	}

	//设置headers
	if len(headers) > 0 {
		for _, header := range headers {
			for k, v := range header {
				req.Header(k, v)
			}
		}
	}
	return req
}

//将interface数据转json
//@param			itf				将数据转成json
//@return							返回json字符串
func InterfaceToJson(itf interface{}) string {
	b, _ := json.Marshal(&itf)
	return string(b)
}

//将interface数据转成int64[如果需要转成int，直接int(int64_number)即可]
//@param			itf				需要转成整型的参数
//@param			num				转换结果
//@param			err				错误
func InterfaceToInt64(itf interface{}) (num int64, err error) {
	return strconv.ParseInt(fmt.Sprintf("%v", itf), 10, 64)
}

//将interface数据转成float64[如果需要转成float32，直接float32(float64_number)即可]
//@param			itf				需要转成整型的参数
//@param			num				转化结果
//@param			err				错误
func InterfaceToFloat64(itf interface{}) (num float64, err error) {
	return strconv.ParseFloat(fmt.Sprintf("%v", itf), 64)
}
