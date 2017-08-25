package golanglocal

import "fmt"

type I interface {
	Mi()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) Mi() {
	fmt.Println(t.S)
}

func InteSatImp() {
	var i I = T{"hello"}
	i.Mi()
}
