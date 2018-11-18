package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Miss Moneypenny"])
	v, ok := fmt.Println(m["Barnabas"])
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v, "is ok")
	}

	m["Todd"] = 34

	for k, v := range m {
		fmt.Printf("Key: %v Value: %v\n", k, v)
	}
}
