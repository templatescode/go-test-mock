package service

import "go-test-mock/calc"

type serviceInterface interface {
	IsMultiplierOfTwo(number int) bool
}

type impl struct {
	serviceInterface
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewService() impl {
	return impl{}
}

// IsMultiplierOfTwo is an implementation
func (s impl) IsMultiplierOfTwo(number int) bool {
	return calc.IsMultiplierOfTwo(number)
}
