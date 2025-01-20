package main

import (
	"fmt"
)

func main() {
	var name string
	var age uint
	quizGame := "Fun Game"
	//Greeting users
	fmt.Printf("Hello, and welcome to %v!\n", quizGame)
	fmt.Println("Type in your name and press Enter: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v!, Let's get started!\n", name)
	fmt.Println("Type in your age and press Enter: ")
	fmt.Scan(&age)

	//Checking if user is old enough to play
	if age >= 18 {
		fmt.Println("Great! You're old enough to play!")
		fmt.Println("Let's get started!")



	} else {
		fmt.Printf("Sorry, you are not old enough to play this game.")
	} else {
		fmt.Println("Invalid input. Please enter a valid age.")
	}

}