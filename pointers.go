package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// pointers
	name := "tifa"
	fmt.Println("memory address of name is: ", &name)
	fmt.Println(name)
	m := &name
	fmt.Println("memory address of m is: ", &m)
	fmt.Println("value at memory address: ", *m)
	updateName(m)
	fmt.Println(name)

}
