package service

type serviceMock struct{}

func newServiceMock() serviceMock {
	return serviceMock{}
}

var mockIsMultiplierOfTwo func(number int) bool

func (s serviceMock) IsMultiplierOfTwo(number int) bool {
	return mockIsMultiplierOfTwo(number)
}
