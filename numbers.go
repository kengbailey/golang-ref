package main

import (
	"fmt"
	"strconv"
)

func main() {

	firstNumber := 1234
	secondNumber := 567890

	fmt.Println("Original numbers:", firstNumber, secondNumber)

	fmt.Println("Can't do this --> string(num)", string(firstNumber), string(secondNumber))

	strNum := strconv.Itoa(firstNumber)
	strNum2 := strconv.Itoa(secondNumber)
	fmt.Println("This works --> strconv.Itoa(num)", strNum, strNum2)

}
