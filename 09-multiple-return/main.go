package main

import (
	"fmt"
	"strings"
)

// Multiple returns: if we want to return more than one value we can use multiple return by using a comma and the type of the values
func getInitial(n string) (string, string) {
	s := strings.ToUpper(n) // converts the string to uppercase
	names := strings.Split(s, " ")

	var initials []string // slice of strings
	for _, v := range names {
		initials = append(initials, v[0:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1] // returning multiple values by using a comma
	} else {
		return initials[0], "_"
	}
}

func main() {

	fn, ln := (getInitial("James Bond"))
	fmt.Println(fn, ln)
}
