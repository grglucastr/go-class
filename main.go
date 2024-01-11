package main

import (
	"fmt"
)

func main() {

	names := []string {"mario", "luigi", "yoshi", "peach"}
	
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

}
