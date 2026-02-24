package main

import "fmt"

func main() {
	age := 30
	name := "Alice"

	//Doesn't add a new line at the end of the output
	fmt.Print("Hello, World! \n")

	//Adds a new line at the end of the output
	fmt.Println("Hello,", name, "You are", age, "years old.")

	//formatted strings allows you to include variables in a string using placeholders

	//%v is a formatspecifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)

	//%q is a formatspecifier that adds double quotes around the value if it's a string
	fmt.Printf("my age is %q and my name is %q \n", age, name)

	//%T is a formatspecifier that prints the type of the variable
	fmt.Printf("age is of type %T \n", age)

	//%f is a formatspecifier that formats a floating-point number
	//you can specify the number of decimal places to display by typing %0.1f where 1 is the number of decimal places you want to display
	fmt.Printf("you scored %f points! \n", 95.5)
	fmt.Printf("you scored %0.1f points! \n", 95.5)
	fmt.Printf("you scored %0.3f points! \n", 95.5)

	//sprintf is a function that returns a formatted string instead of printing it to the console
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("saved string:", str)
}
