package mylib

// SwapStrings swaps two strings and returns them in reverse order
func SwapStrings(a, b string) (string, string) {
	return b, a
}

// AddNumbers adds two integers and returns the result
func AddNumbers(x, y int) int {
	return x + y
}

// IsEven checks if a number is even using the modulus operator and if statement
func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

// NumberSign returns a string indicating if a number is positive, negative, or zero using if-else
func NumberSign(n int) string {
	if n > 0 {
		return "positive"
	} else if n < 0 {
		return "negative"
	}
	return "zero"
}

// DayType uses switch to return if a day is a weekday or weekend
func DayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "unknown"
	}
}

// SumUpTo uses a for loop to sum all numbers up to n (inclusive)
func SumUpTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
