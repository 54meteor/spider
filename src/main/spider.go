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

func (s *Spider) getAPI() {
	for i, v := range s.UrlList {
		s.Chs[i] = make(chan int)
		s.An.Path = s.Path
		go s.An.GetContent(v, s.FileName[i], i, s.Chs[i])
	}
}

func (s *Spider) getHTML(filter string, tag string, fileNameFilter string) {
	dir := new(util.Dir)
	dir.FilePath = s.Path
	fileList := dir.GetFileList()
	for _, file := range fileList {
		cache := new(util.File)
		cache.FilePath = s.Path
		cache.FileName = file.Name()
		content, _ := ioutil.ReadFile(cache.FilePath + cache.FileName)
		cache.Content = string(content)
		value := gjson.Get(cache.Content, tag)
		for key, url := range value.Array() {
			s.Chs[key] = make(chan int)
			s.An.Path = s.Path
			s.An.ConFilter.Grep = filter
			nameFilter := new(util.Filter)
			nameFilter.Grep = fileNameFilter
			nameFilter.Content = url.Str
			filename := nameFilter.Filter()
			go s.An.GetContent(url.Str, filename, key, s.Chs[key])
		}
	}
}
