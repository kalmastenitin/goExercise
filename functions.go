import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v.\n", n)
}

// passing function as argument
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Ramesh")
	cycleNames([]string{"roni", "bonut", "goram"}, sayGreeting)
	fmt.Printf("area of circle is: %0.2f", circleArea(10.5))
}
