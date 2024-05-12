package main

import "fmt"

// Pointers: A pointer is a variable that stores the memory address of another variable
// We get a pointer by using the & operator

// to pass a pointer we need to add an * in front of the variable type
func updateName(n *string) {
	*n = "Link"
}

func main() {

	name := "Zelda"

	// updateName(name)

	// Here we get the memory address of the variable
	// fmt.Println("Memory address of name:", &name)

	m := &name // we declare "m" as a pointer to the variable "name", m has their own memory address and stores the memory address of the variable "name"

	fmt.Println("Memory address of m:", m)

	// if we have a pointer we can dereference it to get the value of the variable at the memory address by using the * operator
	fmt.Println("value at memory address m:", *m)

	// in other words:
	// - We store the memory addres of the variable "name" in a pointer p using the &p operator
	// - We dereference the pointer p to get the value of the variable using the *p operator

	/*
		|--name--|---m----|
		|  0x001 | 0x002  |
		| "Zelda"| p0x001 | // the p is for pointer
		|--------|--------|
	*/

	// we could pass pointers to functions too so we can update the original variable using the pointer instead of overwriting it
	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

	// The copy created we pass by value is pointing to the same memory address than the original variable, for that reason the original value is affected.
}
