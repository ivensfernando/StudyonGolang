package golanglocal

import (
	"fmt"
	//"math"
)

//type Vertex struct {
//	X, Y float64
//}

//func Abs(v VertexMethods) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func Scale(v VertexMethods, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleP(v *VertexMethods, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MetPontExplain() {
	var p = VertexMethods{3, 4}
	//v := VertexMethods{3, 4}
	Scale(p, 10)
	fmt.Println(Abs(p))
// Loock the difference
	v := VertexMethods{3, 4}
	ScaleP(&v, 10)
	fmt.Println(Abs(v))
}