package main

import ("fmt")

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(bulkSend(21))

	fmt.Println(maxMessages(10.0))

	// while/for
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")

	// continue 
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
		  continue
		}
		fmt.Println(i)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
		  break
		}
		fmt.Println(i)
	}
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