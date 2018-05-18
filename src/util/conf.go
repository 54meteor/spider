package util

import (
	"io/ioutil"

	"github.com/tidwall/gjson"
)

type Config struct {
	fileName string
}

func (c *Config) InitConfig(path string) map[string]gjson.Result {
	content, _ := ioutil.ReadFile(path)
	cnt := string(content)
	//使用Json解析库解析json数据
	value := gjson.Get(cnt, "cfg")
	return value.Map()
}
