package main

import (
	"fmt"
)

func main() {
	var name string
	var quizGame string = "Fun Game"
	//Greeting users
	fmt.Printf("Hello, and welcome to %v!\n", quizGame)
	fmt.Println("Type in your name and press Enter: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v!, Let's get started!\n", name)

}