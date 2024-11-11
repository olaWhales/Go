package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello , %s\n", name)

	var age int
	fmt.Print("How old are you? : ")
	fmt.Scanln(&age)
	fmt.Print("I'm ", age , " years old")
}


