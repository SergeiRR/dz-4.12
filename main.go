package main

import (
	"fmt"
)

func main() {

	var userInformation string
	fmt.Print(`_Считывание ввода  виде отдельной функции_ `)
	fmt.Println(userInformation)
	fmt.Scan(&userInformation)

}
