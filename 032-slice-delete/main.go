package main

import "fmt"

func main() {
	x := []string{"John", "James", "Joe"}
	for i, _ := range x {
		fmt.Println(len(x))
		if len(x) > 1 {
			x = append(x[:i], x[i+1:]...)
		}
		fmt.Println("after", x)
	}
}
