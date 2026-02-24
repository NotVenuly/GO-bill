package main

import (
	"fmt"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie"}

	for index, value := range names {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	for _, value := range names {
		fmt.Printf("Value: %v\n", value)
	}

	for index := range names {
		fmt.Printf("Index: %v\n", index)
	}
}
