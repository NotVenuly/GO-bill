package main

import (
	"fmt"
)

func main() {
	age := 30

	fmt.Println(age <= 50)
	fmt.Println(age >= 18)
	fmt.Println(age == 18)
	fmt.Println(age != 50)

	names := []string{"Alice", "Bob", "Charlie"}

	for index, name := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		fmt.Printf("the value at position %v is %v\n", index, name)
	}
}
