package golanglocal

import "fmt"

type VertexM struct {
	Lat, Long float64
}

var m map[string]VertexM

func MpsGo() {
	m = make(map[string]VertexM)
	m["Bell Labs"] = VertexM{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
