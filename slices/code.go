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

	// make
	// func make([]T, len, cap) []T
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)

	// the capacity argument is usually omitted and defaults to the length
	mySlice1 := make([]int, 5)
	fmt.Println(mySlice1)

	mySlice2 := []string{"I", "love", "go"}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	fmt.Println(getMessageCosts(
		[]string{
			"qwertyu",
			"fghddjksdd",
			"gdhjsklaskd",
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

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}