package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	promptOptions()
}
