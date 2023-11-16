package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = appfront(1, digits)
	return digits
}

func appfront(x int, slice []int) []int {
	slice = append([]int{x}, slice...)
	return slice
}

func main() {
	plusOne([]int{1, 2, 3})
}
