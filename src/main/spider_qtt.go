package main

import (
	"domain"
	"fmt"
	"io/ioutil"
	"strconv"
	"util"
	//	"github.com/tidwall/gjson"
)

type Spider_qtt struct {
	Strat, End int
	Path, Id   string
	Chs        []chan int
}

func (s *Spider_qtt) getContentList() *Spider_qtt {
	domain.InitUrlList()
	for _, v := range domain.UrlList {
		for i := s.Strat; i < s.End; i++ {
			s.Chs[i] = make(chan int)
			go s.getList(i, v, s.Chs[i])
		}
	}
	return s
}

func (s *Spider_qtt) getChan() {
	for _, ch := range s.Chs {
		_, ok := <-ch
		if ok {
			close(ch)
		}
	}
}

func (s *Spider_qtt) getList(step int, v string, ch chan int) {
	io, err := util.GetContent(v + "page=" + strconv.Itoa(step+1) + "&cid=" + s.Id)
	if err != nil {
		return
	}
	f := new(domain.File)
	f.FilePath = s.Path
	f.FileName = s.Id + "_" + strconv.Itoa(step)
	f.Content = string(io)
	file, err := f.CreateFile()
	f.F = file
	f.WriteFile()
	ch <- step

}

func (s *Spider_qtt) getContent() {
	f := new(domain.File)
	f.FilePath = s.Path
	fileList := f.GetFileList()
	for _, file := range fileList {
		cache := new(domain.File)
		cache.FilePath = s.Path
		cache.FileName = file.Name()
		content, _ := ioutil.ReadFile(f.FilePath + f.FileName)
		fmt.Println(string(content))
		//		f.Content = string(content)
		//		fmt.Println(f.Content)
	}
	//	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//	value := gjson.Get(json, "age")
	//	fmt.Println(value.String())
}
