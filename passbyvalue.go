package main

import "fmt"

func updateName(x string) string {
	// copy of x is created and stored in memory
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["tea"] = 2.00
}

func main() {
	// pass by value
	// group A variables (int, string, float, bool, arrays, structs)
	name := "rakesh"
	name = updateName(name)
	fmt.Println(name)

	// group B variables (slices, map, functions)
	menu := map[string]float64{
		"pie":  4.44,
		"cake": 5.00,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
