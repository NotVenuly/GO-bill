package main

import "fmt"

func main() {
	var nameOne string = "John"
	var nameTwo = "Doe"
	var nameThree string
	nameFour := "Smith"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	var ageOne int = 30
	var ageTwo = 25
	var ageThree int
	ageFour := 40
	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	//bits and memory
	var bigNumber int8 = 52
	var smallNumber int8 = -128
	var num int16 = 32767
	var mediumNumber uint8 = 25
	fmt.Println(bigNumber, smallNumber, num, mediumNumber)

	var scoreOne float32 = 95.5
	var scoreTwo float64 = 89.75
	scoreThree := 78.25

	println(scoreOne, scoreTwo, scoreThree)
}
