package main

import "fmt"

// reverse the given word

func reverseIt(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}
	return reversed
}

func main() {
	wordList := []string{"Jocker", "cat", "monkey", "nitin"}
	for _, i := range wordList {
		reverse := reverseIt(i)
		fmt.Println("Reverse String : ", reverse)
	}

}
