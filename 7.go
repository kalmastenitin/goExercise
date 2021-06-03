package main

import "fmt"

// sum of two indexes shoud equal to provided sum
func checkSum(list []int, sum int) (int, int) {
	for i, val1 := range list {
		for j, val2 := range list {
			if i == j {
				continue
			}
			if val1+val2 == sum {
				return i, j
			}

		}
	}
	return -1, -1
}

func main() {
	index1, index2 := checkSum([]int{1, 2, 3, 4, 5, 6, 7, 8}, 15)
	fmt.Println("matching indexes are ", index1, index2)
}
