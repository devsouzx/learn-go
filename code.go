package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	variables()
}

func variables() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
