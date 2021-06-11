package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func promptOptions() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		fmt.Println("You Have Selected :a")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
	case "s":
		fmt.Println("You Have Selected :s")
	case "t":
		fmt.Println("You Have Selected :t")
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("You hace selected invalid option!")
	}
}

func main() {
	promptOptions()
}
