package main

import (
	"domain"
	"util"
)

func main() {
	domain.InitUrlList()
	for _, v := range domain.UrlList {
		io, err := util.GetContent(v)
		if err != nil {
			return
		}
		f := new(domain.File)
		f.FilePath = "../../res/"
		f.FileName = "255_65"
		f.Content = string(io)
		file, err := f.CreateFile()
		f.F = file
		f.WriteFile()
	}
}
