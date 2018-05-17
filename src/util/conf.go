package util

import (
	"io/ioutil"

	"github.com/tidwall/gjson"
)

type Config struct {
	fileName string
}

func (c *Config) InitConfig(path string, fileName string) map[string]gjson.Result {
	f := new(File)
	f.FilePath = path
	f.FileName = fileName
	content, _ := ioutil.ReadFile(f.FilePath + f.FileName)
	f.Content = string(content)
	//使用Json解析库解析json数据
	value := gjson.Get(f.Content, "cfg")
	return value.Map()
}
