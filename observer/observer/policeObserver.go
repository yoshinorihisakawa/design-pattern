package observer

import "fmt"

type PoliceObserver struct {
	str string
}

func (p *PoliceObserver) Update(str string) {
	police := PoliceObserver{str:str}
	police.Send()
}

func (p *PoliceObserver) Send() {
	fmt.Println("police : ", p.str)
}
