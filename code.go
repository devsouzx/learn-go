package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	shortVariableDeclaration()
}

func basicVariables() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func shortVariableDeclaration() {
	messageStart :=  "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	
	fmt.Println(messageStart, age, messageEnd)
}
