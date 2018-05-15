package main

import (
	"domain"
	"io/ioutil"
	"regexp"
	"strconv"
	"util"

	"github.com/tidwall/gjson"
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
			go s.getUrl(v+"page="+strconv.Itoa(i+1)+"&cid="+s.Id,
				s.Id+"_"+strconv.Itoa(i), i, s.Chs[i])
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

func (s *Spider_qtt) getUrl(url string, fileName string, key int, ch chan int) {
	io, err := util.GetContent(url)
	if err != nil {
		return
	}
	f := new(domain.File)
	f.FilePath = s.Path + "content/"
	f.FileName = fileName

	//	re, _ := regexp.Compile("\\<!doc[\\S\\s]+?\\<section")
	//	src := re.ReplaceAllString(string(io), "<section")
	re, _ := regexp.Compile("\\<section[\\S\\s]+?\\</section\\>")
	src := re.FindString(string(io))

	//	re, _ = regexp.Compile("</section\\>[\\S\\s]+?\\</html\\>")
	//	src = re.ReplaceAllString(src, "</section>")

	f.Content = src

	file, err := f.CreateFile()
	f.F = file
	f.WriteFile()
	ch <- key
}

func (s *Spider_qtt) getContent() *Spider_qtt {
	f := new(domain.File)
	f.FilePath = s.Path
	fileList := f.GetFileList()
	for _, file := range fileList {
		cache := new(domain.File)
		cache.FilePath = s.Path
		cache.FileName = file.Name()
		content, _ := ioutil.ReadFile(cache.FilePath + cache.FileName)
		cache.Content = string(content)
		value := gjson.Get(cache.Content, "data.data.#.url")
		for key, url := range value.Array() {
			s.Chs[key] = make(chan int)
			go s.getUrl(url.Str, strconv.Itoa(key), key, s.Chs[key])
		}
	}
	return s
}
