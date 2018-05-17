package main

import (
	"strconv"
	"util"
)

//入口程序，目前用于配置抓取的必要信息，后续调整为读取配置文件

func main() {
	//创建爬虫对象
	s := new(Spider)
	//配置起始页码
	s.Strat = 0
	//配置结束页码
	s.End = 10
	//配置存储目录
	s.Path = "../../res/"
	//配置要抓取的栏目id
	s.Id = "255"
	//初始化通道
	s.Chs = make([]chan int, s.End)
	//创建要抓取的地址列表
	for i := s.Strat; i < s.End; i++ {
		//创建参数map
		urls := make(map[string]string)
		urls["page"] = strconv.Itoa(i + 1)
		urls["cid"] = s.Id
		//创建url
		s.UrlList = append(s.UrlList, util.CreateUrlList("http://api.1sapp.com/content/outList?tn=1", urls))
		//创建对应的存储文件名
		s.FileName = append(s.FileName, urls["cid"]+"_"+urls["page"])
	}

	//抓取接口
	//	s.getAPI()
	//抓取html
	//	s.getLocalHTML("\\<section[\\S\\s]+?\\</section\\>", "data.data.#.url", "\\/[\\d\\d]+?\\.html")

	s.getHtml([]string{"http://www.baidu.com"}, "\\<div class=\"s_form\"[\\S\\s]+?\\</div></div>", "baidu.com")
	s.An.GetChan(s.Chs)

}
