import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	// packages in standard library
	// Strings Package
	greeting := "hello everyone"
	// check string in given sentance
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "hi"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(greeting)
	// original value doesnot get changed just it returns value
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "one"))
	fmt.Println(strings.Split(greeting, " "))

	// sort package
	ages := []int{45, 20, 25, 30, 75, 60, 55}
	sort.Ints(ages) // it changes the elemts of original list
	fmt.Println(ages)

	index := sort.SearchInts(ages, 300) // it returns extra index
	fmt.Println(index)

	names := []string{"nitin", "noman", "nomad"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "rohan")) // it returns extra elements
}
