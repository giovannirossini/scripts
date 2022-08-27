package math

// Calculation function is a generic math to any type of calculation
func Calculation(function func(int, int) int, numA int, numB int) int {
	return function(numA, numB)
}
