package golanglocal

import "fmt"

//type Vertex struct {
//	X, Y float64
//}

//func (v *VertexMethods) Scale1(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

func ScaleFunc(v *VertexMethods, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Indirection() {
	v := VertexMethods{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &VertexMethods{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
