package main

import (
	"fmt"
	"time"
)

func main() {
	sendMessage(sendingReport{
		reportName: "First report",
		numberOfSends: 10,
	})
	sendMessage(birthdayMessage{
		recipientName: "Joao",
		birthdayTime: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}