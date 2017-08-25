package golanglocal

import (
	"fmt"
	"math"
)

type Ii interface {
	Mf()
}
//
//type T struct {
//	S string
//}

func (t *T) Mf() {
	fmt.Println(t.S)
}

type F float64

func (f F) Mf() {
	fmt.Println(f)
}

func InterValue() {
	var i Ii

	i = &T{"Hello"}
	describeIV(i)
	i.Mf()

	i = F(math.Pi)
	describeIV(i)
	i.Mf()
}

func describeIV(i Ii) {
	fmt.Printf("(%v, %T)\n", i, i)
}
