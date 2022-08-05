package main

import "fmt"

var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {

	for number := range numbers {

		if number%2 == 0 {
			fmt.Println(number, " - Even")
		} else {
			fmt.Println(number, " - Odd")
		}
	}
}
