package main

import "fmt"

func main() {

	x := func() {
		fmt.Println("This is x")
	}

	foo(x)

}

func foo(x func()) {
	x()
	fmt.Println("This is foo")

}
