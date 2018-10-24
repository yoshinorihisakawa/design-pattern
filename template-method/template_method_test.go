package template_method

import (
	"testing"
	"fmt"
)

func TestNormalDisplay(t *testing.T) {
	expected := "======\n|normal|\n|normal|\n|normal|\n|normal|\n|normal|\n======"
	display := Display{&NormalDisplay{str: "normal"}}
	result := display.Display()
	if result != expected {
		t.Errorf("Expect result to %s, but %s.\n", expected, result)
	}
	fmt.Println("Response",result)
}

func TestGorgeousDisplay(t *testing.T) {
	expected := "＊*＊*＊*＊*＊\n＊gorgeous＊\n＊gorgeous＊\n＊gorgeous＊\n＊gorgeous＊\n＊gorgeous＊\n＊*＊*＊*＊*＊"
	display := Display{&GorgeousDisplay{Str: "gorgeous"}}
	result := display.Display()
	if result != expected {
		t.Errorf("Expect result to \n%s, but \n%s.\n", expected, result)
	}
	fmt.Println("Response",result)
}
