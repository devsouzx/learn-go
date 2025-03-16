package main

import (
	"fmt"
)

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

// Embedded Structs
type truck struct {
	car 
	bedSize int
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

	lanesTruck := truck{
		bedSize: 10,
		car: car{
			brand: "Toyota",
    		model: "Camry",
		},
	}

	fmt.Println(lanesTruck.brand)
	fmt.Println(lanesTruck.model)
	
	fmt.Println(r.area())
	fmt.Println(authInfo.data())

	fmt.Println(newUser("Joao", "basic"))
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

// Struct Methods
func (r rect) area() int {
	return r.width * r.height
}

var r = rect{
	width: 5,
	height: 10,
}

type rect struct {
	width int
	height int
}

type authenticationInfo struct {
	username string
	password string
}

var authInfo = authenticationInfo{
	username: "joao",
	password: "blabla",
}

func (authInfo authenticationInfo) data() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

// update user excercise
type User struct {
	Membership
	Name string
}

type Membership struct {
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

type MessageResponse struct {
	message string
	canSend bool
}

func (user User) sendMessage(messsage string, messageLength int) MessageResponse {
	var messageResponse = MessageResponse{
		message: messsage,
		canSend: true,
	}
	if (messageLength > user.MessageCharLimit) {
		messageResponse.canSend = false
		messageResponse.message = ""
	}
	return messageResponse
}