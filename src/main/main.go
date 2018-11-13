package main

import (
	"flag"
	"fmt"
	"strconv"
	"util"
)

//入口程序，读取配置文件，启动爬虫

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path")
	flag.Parse()
	if len(path) == 0 {
		fmt.Println("invalid path")
		return
	}
	config := new(util.Config)
	cfg := config.InitConfig(path)
	//创建爬虫对象
	s := new(Spider)
	//配置存储目录
	s.Path = cfg["path"].Str
	//初始化通道
	s.Chs = make([]chan int, 10)
	for key, v := range cfg["urls"].Array() {
		s.UrlList = append(s.UrlList, v.Str)
		s.FileName = append(s.FileName, strconv.Itoa(key))
	}
	switch cfg["type"].Str {
	case "api":
		s.getAPI()
	case "local":
		s.getLocalHTML("\\<section[\\S\\s]+?\\</section\\>", "data.data.#.url", "\\/[\\d\\d]+?\\.html")
	case "html":
		s.getHtml([]string{"http://www.baidu.com"}, "\\<div class=\"s_form\"[\\S\\s]+?\\</div></div>", "baidu.com")
	case "xml":
		s.getXML()
	case "decodexml":
		s.getXMLContent()
	}
	//抓取接口
	//	s.getAPI()
	//抓取html
	//	s.getLocalHTML("\\<section[\\S\\s]+?\\</section\\>", "data.data.#.url", "\\/[\\d\\d]+?\\.html")

	//	s.getHtml([]string{"http://www.baidu.com"}, "\\<div class=\"s_form\"[\\S\\s]+?\\</div></div>", "baidu.com")
	s.An.GetChan(s.Chs)
}
