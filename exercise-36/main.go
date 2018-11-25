package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)

	bi, bs := bar()
	fmt.Println(bi)
	fmt.Println(bs)

}

func foo() int {
	return 4
}

func bar() (int, string) {
	return 5, "bar"
}
