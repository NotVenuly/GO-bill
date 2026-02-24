package main

import (
	"fmt"
	"math"
)

func SayGreeting(name string) {
	fmt.Printf("Good morning, %v\n", name)
}

func sayBye(name string) {
	fmt.Printf("Good bye, %v\n", name)
}

func CycleNames(names []string, function func(string)) {
	for _, value := range names {
		function(value)
	}
}

func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	SayGreeting("John")
	sayBye("John")

	CycleNames([]string{"John", "Jane", "Doe"}, SayGreeting)
	CycleNames([]string{"John", "Jane", "Doe"}, sayBye)

	a1 := CircleArea(5)
	a2 := CircleArea(10)

	fmt.Printf("Area 1: %0.2f\n", a1)
	fmt.Printf("Area 2: %0.2f\n", a2)
}
