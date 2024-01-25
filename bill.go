package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

func (b *bill) format() string {

	fs := "Bill breakdown: \n"
	var total float64 = 0
	total += b.tip

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v\t .... \t$%v \n", k+": ", v)
		total += v
	}

	fs += fmt.Sprintf("\n%-25v\t .... \t$%0.2f\n", "Tip: ", b.tip)

	fs += fmt.Sprintf("\n\n%-25v\t .... \t$%0.2f", "Total: ", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Bill saved to file %v\n.", b.name+".txt")
}
