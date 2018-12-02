package main

import "fmt"

func main() {

	f := foo()
	f()

}

func foo() func() {
	fmt.Println("Foo")
	return func() {
		fmt.Println("Bar")
	}
}
