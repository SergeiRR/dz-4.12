package main

import (
	"fmt"
)

func main() {

	userInformation := getUserInput()
	fmt.Println(userInformation)
}
func getUserInput() string {
	var input string
	fmt.Scan(&input)
	return input

}
func convertCurrency(amount float64, fromCurrency string,
	toCurrency string) float64 {
	return 0

}
