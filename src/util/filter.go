package util

import (
	"fmt"
	"regexp"
)

type Filter struct {
	Grep    string
	Content string
}

//使用正则匹配寻找字符串
func (f *Filter) Filter() string {
	re, _ := regexp.Compile(f.Grep)
	src := re.FindString(f.Content)
	fmt.Println(src)
	return src
}
