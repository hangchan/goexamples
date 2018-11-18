package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for i, v := range m {
		fmt.Println(i)
		for ii, vv := range v {
			fmt.Println(ii, vv)
		}
	}

}
