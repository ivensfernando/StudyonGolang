package golanglocal

import (
	"fmt"
	"math"
)

//type MyFloat float64

func (f MyFloat) Absm() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MetCont() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Absm())
}
