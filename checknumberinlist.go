package main

import (
	"fmt"
)

// check list containes the given number or not
func numList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}

func main() {
	if numList([]int{1, 2, 3, 4, 5}, 19) {
		fmt.Println("number is inside list")
	} else {
		fmt.Println("number is absent")
	}
}
