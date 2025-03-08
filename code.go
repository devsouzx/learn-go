package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	sameLineDeclaration()
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

func typeSizes() {
	accountAgeFloat := 2.6
	
	accountAgeInt := int64(accountAgeFloat)
	
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

func sameLineDeclaration() {
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	
	fmt.Println(averageOpenRate, displayMessage)
}
