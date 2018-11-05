package main

import "fmt"

const a string = "foo"
const b = "bar"

func main() {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", b)
	fmt.Println(b)
}
