package math

// Sum is a math function to return the sum of two numbers
func Sum(x int, y int) (result int) {
	result = x + y
	return
}

// Multiplier is a math function to return the product of two numbers
func Multiplier(x int, y int) (result int) {
	result = x * y
	return
}

// Minus is a math function to return the subtraction of two numbers
func Minus(x int, y int) (result int) {
	result = x - y
	return
}

// Divide is a math function to return the division of two numbers
func Divide(x int, y int) (result int) {
	result = x / y
	return
}

// DivideWithRest is the same function of Divide, however with the rest of division
func DivideWithRest(x int, y int) (result int, rest int) {
	result = x / y
	rest = x % y
	return
}
