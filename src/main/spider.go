package main

func main() {
	s := new(Spider_qtt)
	s.Strat = 0
	s.End = 71
	s.Path = "../../res/"
	s.Id = "255"
	s.Chs = make([]chan int, s.End)
	//	s.getContentList().getChan()
	s.getContent().getChan()
}
