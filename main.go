package main

import (
	"fmt"
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

func main() {

	names := []string{"cloud", "tifa", "barret"}
	
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

}
