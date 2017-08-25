package golanglocal

import "fmt"

type VertexP struct {
	X int
	Y int
}

func PointOfStruct() {
	v := VertexP{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

