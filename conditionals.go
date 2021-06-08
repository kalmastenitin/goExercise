package main

import (
	"fmt"
)

func main() {
	// if else
	age := 46
	if age < 20 {
		fmt.Println("age is less than 20")
	} else if age > 20 && age <= 45 {
		fmt.Println("age is between 21 to 45")
	} else {
		fmt.Println("age is greater than 46")
	}

	names := []string{"rakesh", "mukesh", "jayesh", "radhe"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continueing at position: ", index)
			continue
		}
		if value == "jayesh" {
			fmt.Println("Breakning at poisition jayesh")
			break
		}
		fmt.Println("the value at %v is %v", index, value)
	}
}
