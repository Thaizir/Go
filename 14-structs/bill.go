package main

// it is a blueprint for a bill
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill, this would work as the constructor operator in JS/TS when we create a new instance of a class
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
