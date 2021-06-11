package main

// receiver functions with pointers

import (
	"fmt"
)

// receiver functions
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"soup": 4.99,
			"pie":    7.99,
			"pizza":  8.45,
			"coffee": 2.45},
		tip: 0,
	}
	return b
}

// format the bill -receiver function
func (b *bill) format() string {
	//this function is associated with bill struct
	//this function is only be called by bill object
	fs := "bill breakdown: \n"
	var total float64 = 0
	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%v ...$%0.2f \n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%v ...$%0.2f \n", "total:", total+b.tip)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	// pointers will be dereferenced by go automatically
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func main() {
	// struct - Blueprint which describe type of data
	mybill := newBill("nitin's bill")
	fmt.Println(mybill)
	mybill.addItem("sausage", 6.44)
	mybill.updateTip(10)
	fmt.Println(mybill.format())

}
