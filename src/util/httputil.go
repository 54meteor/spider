package util

import (
	"io/ioutil"
	"net/http"
)

//获取目标url的内容
func GetContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	io, err := ioutil.ReadAll(resp.Body)
	return io, err
}
