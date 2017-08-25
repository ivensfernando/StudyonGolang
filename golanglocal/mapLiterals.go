package golanglocal

import "fmt"

type VertexML struct {
	Lat, Long float64
}

var m0 = map[string]VertexML{
	"Bell Labs": VertexML{
		40.68433, -74.39967,
	},
	"Google": VertexML{
		37.42202, -122.08408,
	},
}

func MapLI() {
	fmt.Println(m0)
}