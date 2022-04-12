package service

import "testing"

//goland:noinspection GoBoolExpressions
func TestIsMultiplierOfTwo(t *testing.T) {

	serviceMock := newServiceMock()

	t.Run("test true", func(t *testing.T) {
		mockIsMultiplierOfTwo = func(number int) bool {
			return true
		}

		expected := true
		result := serviceMock.IsMultiplierOfTwo(5)

		if expected != result {
			t.Errorf("\nExpected: %v\nResult: %v\n", expected, result)
		}
	})

}
