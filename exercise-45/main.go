package main

import "fmt"

func main() {
	x := "main"

	foo(x)

}

func foo(x string) {
	x = "foo"
	fmt.Println(x)
}
