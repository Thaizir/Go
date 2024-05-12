package main

import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello!"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll")) // returns the position where starts "ll"
	// fmt.Println(strings.Split(greeting, " ")) // returns an array of strings where each string is a word

	// the original value is unchanges
	// fmt.Println(greeting)

	ages := []int{45, 20, 32, 19, 22, 11, 70, 60, 25}
	sort.Ints(ages) // sorts the ages in ascending order and alter the original array
	fmt.Println(ages)

	index := sort.SearchInts(ages, 19) // if the searched in the array doesn't exits will return the lenght of the array/slice +1
	fmt.Println(index)

	names := []string{"Link", "Zelda", "Sheik", "Ganondorf", "Ruto"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) // seach in the slice the position where bowser is
}
