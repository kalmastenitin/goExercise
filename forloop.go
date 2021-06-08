package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x:", x)
		x++
	}

	for n := 3; n < 6; n++ {
		fmt.Println("value of n:", n)
	}

	names := []string{"john", "mario", "yogi", "santa"}
	for i := 0; i < len(names); i++ {
		fmt.Println("name here is", names[i])
	}
	for index, value := range names {
		fmt.Printf("value at index %v is: %v\n", index, value)

	}

}
