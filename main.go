package main

import "fmt"

func main() {
	var ages = [3]int{25, 30, 35}

	names := [4]string{"Alice", "Bob", "Charlie", "David"}
	names[0] = "Eve"

	scores := []int{90, 85, 92}
	scores = append(scores, 100)

	// gets all the elements from the given number until the given number or the end of the array
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:2]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	fmt.Println(ages, len(ages))

	fmt.Println(names, len(names))

	fmt.Println(scores, len(scores))
}
