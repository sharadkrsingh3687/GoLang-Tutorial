package main

import "fmt"

func main() {

	number := 2
	// check give number is prime or not

	if isPrimeNumber(number) == true {
		fmt.Println(number, " is prime number")
	} else {
		fmt.Println(number, " is not prime number")
	}
}

func isPrimeNumber(number int) bool {

	for index := 2; index < number; index++ {
		if number%index == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}
