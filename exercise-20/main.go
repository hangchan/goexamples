package main

import "fmt"

func main() {

	switch {
	case true:
		{
			fmt.Println("true")
		}
	case (1 == 2):
		{
			fmt.Println("false")
		}
	}
}
