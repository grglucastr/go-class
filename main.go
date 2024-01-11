package main

import (
	"fmt"
)

func main() {

	names := []string {"mario", "luigi", "yoshi", "peach"}
	
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}

}
