package main

import (
	"io/ioutil"
	"util"

	"github.com/tidwall/gjson"
)

type Spider struct {
	Strat, End        int
	Path, Id          string
	Chs               []chan int
	UrlList, FileName []string
	An                util.Analysis
}

//抓取API类接口
func (s *Spider) getAPI() {
	for i, v := range s.UrlList {
		s.Chs[i] = make(chan int)
		s.An.Path = s.Path
		go s.An.GetContent(v, s.FileName[i], i, s.Chs[i])
	}
}

//抓取html类页面
func (s *Spider) getHTML(filter string, tag string, fileNameFilter string) {
	//从本地读取文件列表
	dir := new(util.Dir)
	dir.FilePath = s.Path
	fileList := dir.GetFileList()
	for _, file := range fileList {
		//将本地文件的内容读入到内存
		cache := new(util.File)
		cache.FilePath = s.Path
		cache.FileName = file.Name()
		content, _ := ioutil.ReadFile(cache.FilePath + cache.FileName)
		cache.Content = string(content)
		//使用Json解析库解析json数据
		value := gjson.Get(cache.Content, tag)
		for key, url := range value.Array() {
			s.Chs[key] = make(chan int)
			s.An.Path = s.Path
			s.An.ConFilter.Grep = filter
			//提取存储使用的文件名
			nameFilter := new(util.Filter)
			nameFilter.Grep = fileNameFilter
			nameFilter.Content = url.Str
			filename := nameFilter.Filter()
			go s.An.GetContent(url.Str, filename, key, s.Chs[key])
		}
	}
}
