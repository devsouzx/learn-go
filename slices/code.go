package main

import "fmt"

func main() {
	// arrays
	var myInts [10]int
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(myInts, primes)
	fmt.Println(getMessageWithRetries("tres", "quatro", "cincore"))
}

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	return [3]string{
			primary,
			secondary,
			tertiary,
		}, [3]int{
			len(primary),
			len(secondary),
			len(tertiary),
		}
}