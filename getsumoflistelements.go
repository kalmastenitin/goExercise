package main

import "fmt"

// get sum of all elements
func sumOf(list []int) int {
	sum := 0
	for _, i := range list {
		sum += i
	}
	return sum
}

func main() {
	sum := sumOf([]int{1, 2, 3, 4, -5, -6, 99})
	fmt.Println("Sum of List Element is: ", sum)
}
