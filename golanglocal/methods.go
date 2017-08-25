package golanglocal

import (
	"fmt"
	"math"
)

type VertexMethods struct {
	X, Y float64
}

func (v VertexMethods) Abss() float64 {
	fmt.Println(v.X)
	fmt.Println(v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsLoc() {
	v := VertexMethods{3, 4}
	fmt.Println(v.Abss())
	v1 := VertexMethods{6, 8}
	fmt.Println(v1.Abss())
}
