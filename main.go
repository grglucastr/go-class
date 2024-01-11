package main

import (
	"fmt"
)

func main() {

	names := []string {"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}
