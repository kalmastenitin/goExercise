package main

import "fmt"

// check string is palindrome or not
func reverseIt(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}
	return reversed
}

func checkPalindrome(list []string) {
	for _, i := range list {
		if i == reverseIt(i) {
			fmt.Println("Is palindrome: ", i)
		} else {
			fmt.Println("Is not palindrome: ", i)
		}

	}
}

func main() {
	checkPalindrome([]string{"cat", "tom", "tat", "nitin", "tenet", "moly"})
}
