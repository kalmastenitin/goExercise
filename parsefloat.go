package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // package to convert string to float64
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func (b *bill) format() string {
	//this function is associated with bill struct
	//this function is only be called by bill object
	fs := "bill breakdown: \n"
	var total float64 = 0
	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)
	return fs
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) updateTip(price float64) {
	b.tip = price
}

func (b *bill) save() {
	fmt.Println("printing the data please wait")
	data := []byte(b.format())
	fmt.Println(data)
	err := os.WriteFile("bills\\"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		fmt.Println("You Have Selected :a")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "s":
		fmt.Println("You Have Selected :s")
		b.save()
	case "t":
		fmt.Println("You Have Selected :t")
		tip, _ := getInput("Enter tip amount ($): ", reader)
		p, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.updateTip(p)
		promptOptions(b)
	default:
		fmt.Println("You hace selected invalid option!")
	}
}

func main() {
	b := newBill("test")
	promptOptions(b)
}
