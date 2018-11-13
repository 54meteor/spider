package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"util"

	"github.com/tidwall/gjson"
)

type Spider struct {
	Path              string
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
func (s *Spider) getLocalHTML(filter string, tag string, fileNameFilter string) {
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
		urls := make([]string, 0, 10)
		for _, url := range value.Array() {
			urls = append(urls, url.Str)
		}
		s.getHtml(urls, filter, fileNameFilter)
	}
}
func (s *Spider) getHtml(value []string, filter string, fileNameFilter string) {
	for key, url := range value {
		s.Chs[key] = make(chan int)
		s.An.Path = s.Path
		s.An.ConFilter.Grep = filter
		//提取存储使用的文件名
		nameFilter := new(util.Filter)
		nameFilter.Grep = fileNameFilter
		nameFilter.Content = url
		filename := nameFilter.Filter()
		go s.An.GetContent(url, filename, key, s.Chs[key])
	}
}

func (s *Spider) getXML() {
	for key, url := range s.UrlList {
		s.Chs[key] = make(chan int)
		s.An.Path = s.Path
		s.An.ConFilter.Grep = ""
		//提取存储使用的文件名
		nameFilter := new(util.Filter)
		nameFilter.Grep = ""
		nameFilter.Content = url
		//		filename := nameFilter.Filter()
		go s.An.GetContent(url, "xml", key, s.Chs[key])
	}
}

func (s *Spider) getXMLContent() {
	file, err := os.Open(s.Path + "data")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	fmt.Println(string(data))

	v := SListBucketResult{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	//	fmt.Println(v)
	getRs(s, v)
}

func getRs(s *Spider, v SListBucketResult) {
	var ch chan int
	for key, element := range v.Content {
		s.An.Path = s.Path
		s.An.ConFilter.Grep = ""
		fmt.Println("=================================>")
		fmt.Println("http://archive.bbx.com/" + element.Key)
		s.An.GetContent("http://archive.bbx.com/"+element.Key, getFileName(element.Key), key, ch)
	}
}

func getFileName(url string) string {
	a := strings.Split(url, "/")
	filename := a[len(a)-1]
	return filename
}

type SListBucketResult struct {
	XMLName xml.Name   `xml:"ListBucketResult"`
	Content []SContent `xml:"Contents"`
}

type SContent struct {
	Key string `xml:"Key"`
}
