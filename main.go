package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type

	phonebook := map[int]string{
		2326: "Luigi",
		6523: "Mario",
		9999: "Emergency",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[9999])

	phonebook[2326] = "bowser"

	fmt.Println(phonebook[2326])
}
