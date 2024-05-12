package main

import "fmt"

func main() {

	// We create a struct using the struct keyword
	// We can create a struct with fields (variables) and methods (functions) inside it
	// It is like a mix between a class and an a type object in TypeScript

	myBill := newBill("Link's bill")
	fmt.Println(myBill)

}
