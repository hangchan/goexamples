package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)

	fmt.Println(x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "the value is", v, "the total is", sum)
	}

	fmt.Println(sum)
	return sum
}
