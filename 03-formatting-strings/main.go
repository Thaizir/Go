package main

import "fmt"

func main() {

	age := 32
	name := "Link"

	// Print: doesn't add a new line as Println
	fmt.Print("Hello, ")
	fmt.Print("Go! \n") // with \n we add a new line
	fmt.Print("new line \n")

	// Println
	fmt.Println("Hello Go!") // adds a new new line
	fmt.Println("Bye Go!")
	fmt.Println("My age is", age, "and my name is", name) // This is a way to output variables

	// Formatted strings: is a way to create string with variables embebed inside.
	// Printf
	// we use format specifiers to introduce the variable. In this case, %v is the formar specifier for variables
	fmt.Printf("My age is %v and my name is %v \n", age, name)

	fmt.Printf("My age is %v and my name is %q \n", age, name) // %q puts the variable within " " quotes but doesn't work with integers
	fmt.Printf("age is of type %T \n", age)                    // %T gets the type of the variable
	fmt.Printf("You scored %0.2f points! \n", 98.85)           // %f gets the float variable and %0.nf gets just n decimals points

	// Springf allow us to save the printed string in a variable
	var str = fmt.Sprintf("My age is: %v and my name is %v \n", age, name)

	fmt.Println("The saved string is:", str)
}
