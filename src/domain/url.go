package domain

var UrlList = make([]string, 0, 10)

func InitUrlList() {
	UrlList = append(UrlList, "http://api.1sapp.com/content/outList?tn=1")
}
