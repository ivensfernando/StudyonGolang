package golanglocal

import (
	"fmt"
	"math"
)

type Abser interface {
	AbsMyFloat() float64
}

func Inter() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VertexMethods{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.

//	a = v  WHAT

	fmt.Println(a.AbsMyFloat())
}

type MyFloat float64

func (f MyFloat) AbsMyFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//type Vertex struct {
//	X, Y float64
//}

func (v *VertexMethods) AbsMyFloat() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
