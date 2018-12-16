package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "James"
	p.last = "Jones"
}

func main() {
	p := person{
		first: "John",
		last:  "Doe",
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}
