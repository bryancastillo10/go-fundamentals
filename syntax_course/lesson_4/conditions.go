package main

import (
	"fmt"
)

// Boolean Conditions /Logical Operations
func main(){
	fmt.Println("Conditional Statements")
	x := 5
	val := x == 5
	fmt.Println("First condition",val)

	new_val := (true || false) && !true
	fmt.Println("Second condition",new_val)
}