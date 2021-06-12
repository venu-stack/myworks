package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("please enter your age : ")

	fmt.Scanln(&age)

	if age > 35 {
		fmt.Println("your good to go ")
	} else if age < 30 {
		fmt.Println("stay home")
	} else {
		fmt.Println("wait")
	}
}
