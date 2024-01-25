package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
}

func (b bill) format() string {

	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v\t .... \t$%v \n", k+": ", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v\t .... \t$%0.2f", "Total: ", total)

	return fs
}
