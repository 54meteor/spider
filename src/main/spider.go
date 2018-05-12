package main

import (
	"domain"
	"os"
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
		f.FileName = "255_65" + ".txt"
		file, err := os.Create("../" + f.FileName)
		f.FilePath = "../"
		f.Content = string(io)
		f.F = file
		f.WriteFile()
	}
}
