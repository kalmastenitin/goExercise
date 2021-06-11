package main

import (
	"fmt"
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

func main() {
	// struct - Blueprint which describe type of data
	mybill := newBill("nitin's bill")
	fmt.Println(mybill)

}
