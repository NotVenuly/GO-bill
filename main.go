package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello, World!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{25, 30, 35, 40, 70, 80, 90}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 35)
	fmt.Println(index)

	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Bob"))
}
