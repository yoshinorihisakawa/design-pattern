package main

import (
	"fmt"
	"github.com/yoshinorihisakawa/design-pattern/observer/observer"
	"github.com/yoshinorihisakawa/design-pattern/observer/subject"
	"math/rand"
	"time"
)

func main() {

	police := &observer.PoliceObserver{}
	swat := &observer.SwatObserver{}

	errSub := subject.ErrorSubject{}

	errSub.RegisterObserver(police)
	errSub.RegisterObserver(swat)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(10)
		if randNum == 7 {
			errSub.DetectError("異常を検知")
		} else {
			fmt.Println("荷物No.:", randNum, " は問題なし.")
		}
		time.Sleep(1 * time.Second)
	}
}
