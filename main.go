package main

import (
	"fmt"
)

func main() {
	name := "Input your name: "
	fmt.Scan(&name)
	fmt.Printf("Hello, %v Welcome to take the quiz game!", name)
	
}