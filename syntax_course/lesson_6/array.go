package main

import "fmt"

// Array and Printing once the similar element
func main() {
	var array []string = []string {"Burger","Fries","Tapsilog","Siopao","Sisig","Palabok","Fries","Cheesecake","Tapsilog","Pancit"}

	for i, element := range array {
		for j := i + 1; j < len(array); j++ {
			element2 := array[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}