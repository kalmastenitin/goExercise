package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":   4.99,
		"pie":    7.99,
		"pizza":  8.45,
		"coffee": 2.45,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as keyTypes
	phonebook := map[int]string{
		2484651: "ramesh",
		3269848: "suresh",
		1598499: "rakesh",
	}
	fmt.Println(phonebook)

	// update value in map
	phonebook[1598499] = "guest"
	fmt.Println(phonebook)
}
