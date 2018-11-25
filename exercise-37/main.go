package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(foo(x...))

	fmt.Println(bar(x))

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}
