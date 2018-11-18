package main

import "fmt"

func main() {
	xxs := [][]string{}
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	xxs = [][]string{xs1, xs2}

	fmt.Println(xs1)
	fmt.Println(xs2)
	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Println("Position", i)
		for j, val := range v {
			fmt.Printf("\t position %v \t value %v\n", j, val)
		}
	}

}
