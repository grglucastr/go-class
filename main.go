package main

import "fmt"

func main() {
	var ages[3] int = [3] int {20, 25, 32}
	var agess = [3]int {20, 25, 32}

	names := [4]string {"yoshi", "mario", "peach", "bowser"}
	names[0] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(agess, len(agess))
	fmt.Println(names, len(names))

	// SLICES
	
	var scores = []int {100, 50, 10}
	scores[2] = 25

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

}
