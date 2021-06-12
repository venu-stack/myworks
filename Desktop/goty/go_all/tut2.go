package main

import (
	"fmt"
)

func main() {
	var age int

	fmt.Println("Enter Your age: ")

	fmt.Scanln(&age)

	if age > 0 {

		fmt.Println("age positive ")
	}
}
