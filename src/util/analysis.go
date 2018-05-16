package util

type Analysis struct {
	Path      string
	ConFilter Filter
}

//解析抓取到的内容，并存储到文件
func (s *Analysis) GetContent(url string, fileName string, key int, ch chan int) {
	//抓取内容
	io, err := GetContent(url)
	if err != nil {
		return
	}
	//创建存储文件
	f := new(File)
	f.FilePath = s.Path
	f.FileName = fileName
	//匹配规则判断。grep有内容，则根据内容进行提炼，如果没有内容，直接将全文赋值
	if len(s.ConFilter.Grep) != 0 {
		s.ConFilter.Content = string(io)
		f.Content = s.ConFilter.Filter()
	} else {
		f.Content = string(io)
	}
	//创建文件
	file, err := f.CreateFile()
	f.F = file
	//保存抓取的内容
	f.WriteFile()
	ch <- key
}

//遍历通道集合
func (s *Analysis) GetChan(chs []chan int) {
	for _, ch := range chs {
		_, ok := <-ch
		if ok {
			close(ch)
		}
	}
}
