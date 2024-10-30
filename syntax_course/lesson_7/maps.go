package main

import (
	"fmt"
)

func main(){
	var mp map[string]int = map[string]int {
		"apple":6,
		"orange":10,
		"pineapple":8,
		"mango":3,
	}

	// mp := make(map[string]int)
	// mp["banana"]=32
	// mp["orange"]=30
	// mp["apple"]=10

	val, bool := mp["app"]
	fmt.Println(val,bool)
	
	fmt.Println(mp)
}