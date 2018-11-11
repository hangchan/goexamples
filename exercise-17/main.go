package main

import "fmt"

func main() {
	bd := 1981

	for {
		if bd > 2018 {
			break
		} else {
			fmt.Println(bd)
			bd++
		}
	}
}
