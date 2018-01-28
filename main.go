package main

import (
	"fmt"

	"github.com/TruthHun/gotil/util"
)

func main() {
	//f := "https://cover.kancloud.cn/johng/gf!middle"
	//f := "https://pic2.zhimg.com/v2-d6010fa144e5647c38ebfe6a93fd619f_b.jpg"
	f := "https://github.com/TruthHun/gotil/archive/master.zip"
	fmt.Println(util.CrawlFile(f, "./", 10))
}
