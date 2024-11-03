package integers

import "fmt"

// Taking two integers and returns the sum of them
func Sum(x, y int) int {
	return x + y
}

func AddIntegers(){
	sum := Sum(1,5)
	fmt.Println(sum)
}

