package main

import "fmt"

func main() {

	           //key   //value
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,		
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps

	for k, v := range menu {
		fmt.Printf("\nKey: %v - Value: %v\n", k, v)
	}

	fmt.Printf("\n------ phonebook ------\n\n")

	// ints as key type
	phonebook := map[int]string {
		5550123: "Jessica",
		5550012: "Kenny",
		5551112: "Lizzy",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[5550012])

	phonebook[5551112] = "Michele"
	fmt.Println(phonebook)

	fmt.Printf("\n------ city population ------\n\n")

	population := map[string]int64{}

	population["New York"] = 20000000
	population["Sao Paulo"] = 11000000
	
	fmt.Println(population)
	
	

}
