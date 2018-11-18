package main

import "fmt"

func main() {

	p1 := struct {
		cheese    bool
		pepperoni bool
		sausage   bool
	}{
		cheese:    true,
		pepperoni: true,
		sausage:   false,
	}

	fmt.Println(p1)

}
