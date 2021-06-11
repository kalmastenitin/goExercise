package main

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
func (b bill) format() string {
	//this function is associated with bill struct
	//this function is only be called by bill object
	fs := "bill breakdown: \n"
	var total float64 = 0
	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}
	// total
	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)
	return fs
}

func main() {
	// struct - Blueprint which describe type of data
	mybill := newBill("nitin's bill")
	fmt.Println(mybill)
	fmt.Println(mybill.format())

}
