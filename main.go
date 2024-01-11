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

	// SLICE RANGES

	rangeOne := names[1:3]
	fmt.Println(rangeOne)

	rangeTwo := names[2:]
	fmt.Println(rangeTwo)

	rangeThree := names[:3] //do not include the the third one
	fmt.Println(rangeThree)

}
