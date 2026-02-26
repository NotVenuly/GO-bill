package main

import "fmt"

func UpdateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	m := &name

	fmt.Println("memory address:", m)
	fmt.Println("value at memory adress:", *m)

	fmt.Println(name)
	UpdateName(m)
	fmt.Println(name)
}
