package util

import (
	"io/ioutil"
	"os"
)

type Dir struct {
	FilePath string
}

func (d *Dir) GetFileList() []os.FileInfo {
	fileList, _ := ioutil.ReadDir(d.FilePath)
	return fileList
}

func (d *Dir) IsDir() bool {
	dir, err := os.Stat(d.FilePath)
	if err != nil {
		return false
	}
	return dir.IsDir()
}

func (d *Dir) CreateDir() bool {
	err := os.MkdirAll(d.FilePath, 0777)
	if err != nil {
		return false
	}
	return true
}
