package template_method_2

import (
	"testing"
	"fmt"
)

func TestNormalDisplay(t *testing.T) {
	expected := "======\n|normal|\n|normal|\n|normal|\n|normal|\n|normal|\n======"
	display := &NormalDisplay{str: "normal"}
	result := display.Display(display)
	if result != expected {
		t.Errorf("Expect result to %s, but %s.\n", expected, result)
	}
	fmt.Println("Response",result)
}

func TestGorgeousDisplay(t *testing.T) {
	expected := "＊*＊*＊*＊*＊\n＊gorgeous＊\n＊gorgeous＊\n＊gorgeous＊\n＊gorgeous＊\n＊gorgeous＊\n＊*＊*＊*＊*＊"
	display := &GorgeousDisplay{Str: "gorgeous"}
	result := display.Display(display)
	if result != expected {
		t.Errorf("Expect result to \n%s, but \n%s.\n", expected, result)
	}
	fmt.Println("Response",result)
}
