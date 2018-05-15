package util

var UrlList = []string{"http://api.1sapp.com/content/outList?tn=1"}

func CreateUrlList(baseUrl string, pm map[string]string) string {
	pms := ""
	for key, value := range pm {
		pms += "&" + key + "=" + value
	}
	return baseUrl + pms
}
