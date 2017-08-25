package golanglocal

//import "fmt"

type Inv interface {
	M()
}

func NilInterVal() {
	var i Inv
	describe(i)
	i.M()
}

//func describeInv(i Inv) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
