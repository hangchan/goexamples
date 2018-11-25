package main

import "fmt"

func main() {
	defer bar()
	foo()
}

func foo() {
	x := "This is foo"
	fmt.Println(x)
}

func bar() {
	x := "This is bar"
	fmt.Println(x)
}
