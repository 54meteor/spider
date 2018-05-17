package util

var UrlList = []string{"http://api.1sapp.com/content/outList?tn=1"}

//组合URL地址，使用baseUrl连接参数
func CreateUrlList(baseUrl string, pm map[string]string) string {
	pms := ""
	//将map的key,value转换成字符串
	for key, value := range pm {
		pms += "&" + key + "=" + value
	}
	return baseUrl + pms
}
