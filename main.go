package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string){
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string){
	fmt.Printf("Goodbye %v\n", name)
}

func cycleNames(names[] string, f func(string)){
	for _, value := range names{
		f(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {

	names := []string{"cloud", "tifa", "barret"}
	
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 3 is %0.3f\n", a1, a2)


}
