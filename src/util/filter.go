package util

import (
	"regexp"
)

type Filter struct {
	Grep    string
	Content string
}

func (f *Filter) Filter() string {
	re, _ := regexp.Compile(f.Grep)
	src := re.FindString(f.Content)
	return src
}
