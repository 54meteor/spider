package domain

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

func (f *File) CheckFileIsExist() bool {
	if _, err := os.Stat(f.FilePath + f.FileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f *File) WriteFile() (int, error) {
	return io.WriteString(f.F, f.Content)
}

func (f *File) CreateFile() {
	if f.CheckFileIsExist() {
		return
	}

}
