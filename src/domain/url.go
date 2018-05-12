package domain

var UrlList = make([]string, 0, 10)

func InitUrlList() {
	UrlList = append(UrlList, "http://api.1sapp.com/content/outList?cid=255&tn=1&page=65&limit=10")
}
