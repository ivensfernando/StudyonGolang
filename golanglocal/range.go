package golanglocal

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func RangeTest() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
