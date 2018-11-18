package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"vanilla", "chocolate", "pecan"},
	}
	p2 := person{
		first:   "Miss",
		last:    "Moneypenny",
		flavors: []string{"chocolate", "raspberry", "lime"},
	}

	fmt.Println(p1)
	for i, v := range p1.flavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2)
	for i, v := range p2.flavors {
		fmt.Println(i, v)
	}

}
