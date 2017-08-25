package exercises
//https://tour.golang.org/moretypes/26
import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	//num := 0
	//value1 := 0
	//value2 := 1
	var num, value1, value2 int = 0 , 0 ,1
	return func () int {
		defer func() { num ++ }()
		if num == 0 {
			//value1 = 1
			//value2 = 0
			//num ++
			return 0
		}//else{
			//if num == 1{
			//	value1 = 1
			//	value2 = 0
			//	//num ++
			//	return 1
			//}else{
				fibonacciNumber := value1 + value2
				value2 = value1
				value1 = fibonacciNumber
				//num ++
				return fibonacciNumber
			//}

		//}

	}


}

func Efibonacci() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(i, f())
	}
}
