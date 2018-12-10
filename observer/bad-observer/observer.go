package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	swatObserver := &SwatObserver{}
	policeObserver := &PoliceObserver{}
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(10)
		if randNum == 7 {
			swatObserver.Update("異常を検知")
			policeObserver.Update("異常を検知")
		} else {
			fmt.Println("荷物No.:", randNum, " は問題なし.")
		}
		time.Sleep(1 * time.Second)
	}
}

type SwatObserver struct {
	str string
}

func (s *SwatObserver) Update(str string) {
	swat := SwatObserver{str: str}
	swat.Send()
}

func (s *SwatObserver) Send() {
	fmt.Println("swat : ", s.str)
}

type PoliceObserver struct {
	str string
}

func (p *PoliceObserver) Update(str string) {
	police := PoliceObserver{str: str}
	police.Send()
}

func (p *PoliceObserver) Send() {
	fmt.Println("police : ", p.str)
}
