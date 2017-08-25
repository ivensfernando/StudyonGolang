package golanglocal

import (
	"fmt"
	"time"
)

func ase() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main2() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println(ase())

	ix := c()
	fmt.Println(ix)

}


