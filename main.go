package main

import "fmt"

func main() {
	var ages[3] int = [3] int {20, 25, 32}
	var agess = [3]int {20, 25, 32}

	names := [4]string {"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(agess, len(agess))
	fmt.Println(names, len(names))

}
