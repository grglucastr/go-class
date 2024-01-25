package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string){

	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1{
		return initials[0], initials[1]
	}

	return initials[0], "_"
}


func main() {

	fn, sn := getInitials("george nunes")
	fmt.Printf("\nFirst name: %v | Last name: %v\n", fn, sn)

	fn2, sn2 := getInitials("lucas bentes")
	fmt.Printf("\nFirst name: %v | Last name: %v\n", fn2, sn2)

	fn3, sn3 := getInitials("gus")
	fmt.Printf("\nFirst name: %v | Last name: %v\n", fn3, sn3)

}
