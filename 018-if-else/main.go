package main

import "fmt"

func main() {
	x := 41
	if x == 40 {
		fmt.Println("This is 40")
	} else if x == 41 {
		fmt.Println("This is 41")
	} else {
		fmt.Println("This is not 40")
	}
}
