package main

import "fmt"

func outFn (newFunc func(int) int){
	fmt.Println(newFunc(7))
}


func main(){
	insideFn := func(x int) int {
		return x +2
	}

	outFn(insideFn)
}