package util

import (
	"io"
	"os"
)

type File struct {
	FileName string
	FilePath string
	Content  string
	F        *os.File
}

//检查文件是否存在
func (f *File) CheckFileIsExist() bool {
	if _, err := os.Stat(f.FilePath + f.FileName); os.IsNotExist(err) {
		return false
	}
	return true
}

//向文件写入内容
func (f *File) WriteFile() (int, error) {
	return io.WriteString(f.F, f.Content)
}

//创建文件
func (f *File) CreateFile() (*os.File, error) {
	dir := new(Dir)
	dir.FilePath = f.FilePath
	//如果目录不存在 ，则创建目录
	if !dir.IsDir() {
		dir.CreateDir()
	}
	if f.CheckFileIsExist() {
		return os.Open(f.FilePath + f.FileName)
	} else {
		return os.Create(f.FilePath + f.FileName)
	}
}
