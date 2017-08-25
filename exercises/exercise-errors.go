package exercises
//https://tour.golang.org/methods/20

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func ExerError() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
