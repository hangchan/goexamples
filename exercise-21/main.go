package main

import "fmt"

func main() {
	favSport := "baseball"
	switch favSport {
	case "basketball":
		{
			fmt.Println("basketball")
		}
	case "football":
		{
			fmt.Println("football")
		}
	case "soccer":
		{
			fmt.Println("soccer")
		}
	case "baseball":
		{
			fmt.Println("baseball")
		}
	default:
		{
			fmt.Println("hockey")
		}
	}
}
