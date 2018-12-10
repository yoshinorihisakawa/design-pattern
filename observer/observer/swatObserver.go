package observer

import "fmt"

type SwatObserver struct {
	str string
}

func (s *SwatObserver) Update(str string) {
	swat := SwatObserver{str:str}
	swat.Send()
}

func (s *SwatObserver) Send() {
	fmt.Println("swat : ", s.str)
}