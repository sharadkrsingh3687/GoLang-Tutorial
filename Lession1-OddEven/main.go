package main

import "fmt"

var numbers []int

func main() {

	var prop bool
	numbers = generateIntSlice(0, 100)

	for _, number := range numbers {
		prop = isOddNumber(number)

		if prop == true {
			fmt.Println(number, " is Odd")
		} else {
			fmt.Println(number, " is Even")
		}
	}

}

func generateIntSlice(rangeStart int, rangeEnd int) []int {

	s := []int{}

	for i := rangeStart; i <= rangeEnd; i++ {
		s = append(s, i)
	}

	return s
}

func isOddNumber(number int) bool {

	if number%2 == 0 {
		return false
	} else {
		return true
	}
}
