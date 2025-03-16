package main

import "fmt"

type car struct {
	brand      string
	model      string
	doors      int
	miliage    int
	frontWheel wheel
	backWheel  wheel
	wheel struct {
		radius   int
		material string
	}
}

type wheel struct {
	radius   int
	material string
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func main() {
	message := messageToSend{}
	message.sender.name = "dgg"
	message.sender.number = 4
	message.recipient.number = 8
	fmt.Println(canSendMessage(message))

	// Anonymous Structs
	myCar := struct {
		brand string
		model string
	} {
		brand: "Toyota",
		model: "Camry",
	}
	fmt.Println(myCar)
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name != "" {
		if mToSend.sender.number > 0 {
			if mToSend.recipient.name != "" {
				if mToSend.recipient.number > 0 {
					return true
				}
			}
		}
	}
	return false
}