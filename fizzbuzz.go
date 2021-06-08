package main

import "fmt"

// check number is divisible by 3, 5

func checkBuzz(list []int) {
	for _, i := range list {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		}
	}
}

func main() {
	checkBuzz([]int{1, 2, 3, 4, 15, 6, 7, 8, 9, 30, 60})
}
