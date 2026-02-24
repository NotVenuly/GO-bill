package main

import (
	"fmt"
	"strings"
)

func GetInitials(name string) (string, string, string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1], initials[2], initials[3]
	}

	return initials[0], "_", "_", "_"
}

func main() {
	firstName, secondName, thirdName, fourthName := GetInitials("John Edgar Allen Doe")

	fmt.Println(firstName, secondName, thirdName, fourthName)
}
