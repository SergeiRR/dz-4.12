package main

import "fmt"

func main() {
	userInput := getUserInput()
	fmt.Println(userInput)
}

func getUserInput() string {
	var input string
	fmt.Print("Введите данные: ")
	fmt.Scan(&input)
	return input
}

func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	return 0
}
