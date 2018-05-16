package util

import (
	"io/ioutil"
	"os"
)

type Dir struct {
	FilePath string
}

//获取目录下文件列表
func (d *Dir) GetFileList() []os.FileInfo {
	fileList, _ := ioutil.ReadDir(d.FilePath)
	return fileList
}

//判断是否是目录
func (d *Dir) IsDir() bool {
	dir, err := os.Stat(d.FilePath)
	if err != nil {
		return false
	}
	return dir.IsDir()
}

//创建目录
func (d *Dir) CreateDir() bool {
	err := os.MkdirAll(d.FilePath, 0777)
	if err != nil {
		return false
	}
	return true
}
