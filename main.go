package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n') // read until the user press enter
	return strings.TrimSpace(input), err
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Bill created.")

	return b
}

func promptOptions(b *bill) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)
		break
	case "s":
		fmt.Println("Save Bill")
		break
	case "t":

		tip, _ := getInput("Enter tip amount ($): ",reader)
		fmt.Println(tip)
		
		break
	default:
		fmt.Println("Invalid option.")
		promptOptions(b)
		break
	}

}

func main() {

	bill := createBill()
	promptOptions(&bill)

	//fmt.Println(bill.format())

}
