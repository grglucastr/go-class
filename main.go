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

func main() {
	sayGreeting("mario")
	sayGreeting("luigi")

	sayBye("mario")
}
