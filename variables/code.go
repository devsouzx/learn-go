package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")

	charactersAndStrings()
}

func basicVariables() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func shortVariableDeclaration() {
	messageStart := "Happy birthday! You are now"
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

func constants() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}

func computedConstants() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)
}

func formatedString() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)

	fmt.Print(msg)
}

func charactersAndStrings() {
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
