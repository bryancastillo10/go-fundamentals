package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	const prefix = "Hello , "
	fmt.Println("Please type your name")
	text := bufio.NewScanner(os.Stdin);
	text.Scan()
	input := text.Text()
	fmt.Printf(prefix + input)
	
}