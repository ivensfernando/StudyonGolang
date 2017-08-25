package golanglocal

import "fmt"

func SliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice2(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice2(s)

	// Extend its length.
	s = s[:4]
	printSlice2(s)

	s = s[:2]
	printSlice2(s)

	s = s[:6]
	printSlice2(s)

	// Drop its first two values.
	s = s[2:]
	printSlice2(s)


}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
