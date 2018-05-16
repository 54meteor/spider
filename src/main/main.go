package main

import (
	"strconv"
	"util"
)

func main() {
	s := new(Spider)
	s.Strat = 0
	s.End = 10
	s.Path = "../../res/"
	s.Id = "255"
	s.Chs = make([]chan int, s.End)

	for i := s.Strat; i < s.End; i++ {
		urls := make(map[string]string)
		urls["page"] = strconv.Itoa(i + 1)
		urls["cid"] = s.Id
		s.UrlList = append(s.UrlList, util.CreateUrlList("http://api.1sapp.com/content/outList?tn=1", urls))
		s.FileName = append(s.FileName, urls["cid"]+"_"+urls["page"])
	}
	s.getAPI()
	//	s.getHTML("\\<section[\\S\\s]+?\\</section\\>", "data.data.#.url", "\\/[\\d\\d]+?\\.html")
	s.An.GetChan(s.Chs)

}
