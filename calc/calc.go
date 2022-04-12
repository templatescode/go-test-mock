package calc

import "fmt"

func IsMultiplierOfTwo(number int) bool {

	fmt.Println("*** without mock")

	if result := number % 2; result != 0 {
		return false
	}
	return true
}
