package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	e := func(x int) {
		fmt.Println("the year big brother started", x)
	}

	e(1984)

}
