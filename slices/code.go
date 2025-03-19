package main

import (
	"errors"
	"fmt"
)

func main() {
	// arrays
	var myInts [10]int
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(myInts, primes)
	fmt.Println(getMessageWithRetries("tres", "quatro", "cincore"))

	// slices
	mySlices := primes[2:4]
	fmt.Println(mySlices)
	lowIndexSlice := primes[3:]
	fmt.Println(lowIndexSlice)
	highIndexSlice := primes[:3]
	fmt.Println(highIndexSlice)
	fmt.Println(primes[:])

	fmt.Println(getMessageWithRetriesForPlan(
		"free",
		[3]string{
			"sdgsgdvgd",
			"hhsgdsdgs",
			"sdhgghsvdg",
		},
	))
	fmt.Println(getMessageWithRetriesForPlan(
		"pro",
		[3]string{
			"sdgsgdvgd",
			"hhsgdsdgs",
			"sdhgghsvdg",
		},
	))
	fmt.Println(getMessageWithRetriesForPlan(
		"other",
		[3]string{
			"sdgsgdvgd",
			"hhsgdsdgs",
			"sdhgghsvdg",
		},
	))
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

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	}
	if plan == planFree {
		return messages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}