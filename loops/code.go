package main

import ("fmt")

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(bulkSend(21))
}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}