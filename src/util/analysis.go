package util

type Analysis struct {
	Path string
}

func (s *Analysis) GetContent(url string, fileName string, key int, ch chan int) {
	io, err := GetContent(url)
	if err != nil {
		return
	}
	f := new(File)
	f.FilePath = s.Path
	f.FileName = fileName
	//	re, _ := regexp.Compile("\\<section[\\S\\s]+?\\</section\\>")
	//	src := re.FindString(string(io))
	f.Content = string(io)

	file, err := f.CreateFile()
	f.F = file
	f.WriteFile()
	ch <- key
}

func (s *Analysis) GetChan(chs []chan int) {
	for _, ch := range chs {
		_, ok := <-ch
		if ok {
			close(ch)
		}
	}
}
