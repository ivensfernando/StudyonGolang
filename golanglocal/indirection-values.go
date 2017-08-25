package golanglocal

import (
	"fmt"
	"math"
)

//type Vertex struct {
//	X, Y float64
//}

//func (v VertexMethods) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func AbsFunc(v VertexMethods) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func IndirectionValues() {
	v := VertexMethods{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &VertexMethods{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
