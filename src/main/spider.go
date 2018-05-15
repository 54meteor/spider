package main

import (
	"io/ioutil"
	"strconv"
	"util"

	"github.com/tidwall/gjson"
)

type Spider struct {
	Strat, End        int
	Path, Id          string
	Chs               []chan int
	UrlList, FileName []string
}

func (s *Spider) getAPI() {
	for i, v := range s.UrlList {
		s.Chs[i] = make(chan int)
		analysis := new(util.Analysis)
		analysis.Path = s.Path
		go analysis.GetContent(v, s.FileName[i], i, s.Chs[i])
	}
}

func (s *Spider) getHTML() {
	dir := new(util.Dir)
	dir.FilePath = s.Path
	fileList := dir.GetFileList()
	for _, file := range fileList {
		cache := new(util.File)
		cache.FilePath = s.Path
		cache.FileName = file.Name()
		content, _ := ioutil.ReadFile(cache.FilePath + cache.FileName)
		cache.Content = string(content)
		value := gjson.Get(cache.Content, "data.data.#.url")
		for key, url := range value.Array() {
			s.Chs[key] = make(chan int)
			analysis := new(util.Analysis)
			analysis.Path = s.Path
			go analysis.GetContent(url.Str, strconv.Itoa(key), key, s.Chs[key])
		}
	}
}
