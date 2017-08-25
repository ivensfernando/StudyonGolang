package golanglocal

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FuncValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	hpotys := func(p, t float64) float64{
		return p * t
	}


	fmt.Println(hypot(5, 12))

	fmt.Println(hpotys(5, 12))

	fmt.Println(compute(hypot))

	fmt.Println(compute(hpotys))

	fmt.Println(compute(math.Pow))
}

