package main

import (
	"fmt"

	"go-test-mock/service"
)

func main() {

	s := service.NewService()
	result := s.IsMultiplierOfTwo(9)

	fmt.Println(result)

}
