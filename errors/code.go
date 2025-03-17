package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sendSMSToCouple("rere", "eeeee"))
	fmt.Println(divide2(6, 2))
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse , err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	return costForCustomer + costForSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func divide2(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
