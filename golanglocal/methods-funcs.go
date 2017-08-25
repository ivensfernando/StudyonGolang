package golanglocal

import (
	"fmt"
	"math"
)

//type VertexMF struct {
//	X, Y float64
//}

func Abs(v VertexMethods) float64 {
	fmt.Println("The X value ",v.X)
	fmt.Println("The Y value ",v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsFuncLoc() {
	v := VertexMethods{3, 4}
	fmt.Println(Abs(v))
}
