package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 1

	if a == b {
		fmt.Println("a = b")
	} else if a == c {
		fmt.Println("a = c")
	} else {
		fmt.Println("a != b")
	}
}
