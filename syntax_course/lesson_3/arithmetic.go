package main

import (
	"fmt"
	"math"
)

func main(){
	circ := math.Pi
	var num1 int = 9
	var num2 int  = 4

	answer := num1%num2
	// For float32
	// fmt.Printf("%g",answer)
	fmt.Printf("%d\n",answer)
	fmt.Println("The value of pi is", circ)

}