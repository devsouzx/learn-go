package main

import ("fmt")

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(bulkSend(21))

	fmt.Println(maxMessages(10.0))
}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}