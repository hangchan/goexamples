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

	m := map[string]person{}

	m[p1.last] = p1
	m[p2.last] = p2

	for _, v := range m {
		// fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for ii, vv := range v.flavors {
			fmt.Println(ii, vv)
		}
	}

}
