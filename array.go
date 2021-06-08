package main

import "fmt"

func main() {
	// declare array here
	var ages = [3]int{20, 30, 25}
	fmt.Println(ages, len(ages))
	var names = [4]string{"nitin", "rakesh", "yogesh", "john"}
	fmt.Println(names)

	//array methods append
	var scores = []int{100, 200, 300}
	fmt.Println(scores[2])
	scores[2] = 25
	fmt.Println(scores)
	newscores := append(scores, 85)
	scores = append(scores, 85)
	fmt.Println(newscores, scores)

	// array slices
	rangeone := scores[1:3]
	fmt.Println(rangeone)
	rangetwo := scores[1:]
	fmt.Println(rangetwo)
	rangethree := scores[:3]
	fmt.Println(rangethree)

	rangeone = append(rangeone, 255)
	fmt.Println(rangeone)
}
