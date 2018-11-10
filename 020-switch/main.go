package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (2 == 2):
		fmt.Println("This should print")
		fallthrough
	case (3 == 3):
		fmt.Println("Will this print?")
		fallthrough
	case (7 == 9):
		fmt.Println("This is false")
	default:
		fmt.Println("This is default")
	}
}
