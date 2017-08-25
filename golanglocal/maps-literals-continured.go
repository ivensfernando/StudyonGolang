package golanglocal

import "fmt"


var m1 = map[string]VertexML{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408}, //order by first element
}

func MLC() {
	fmt.Println(m1)
}
