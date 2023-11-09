package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := splitDigits(x)
	reverse := reverseDigits(digits)
	return compareList(digits, reverse)
}

func compareList(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func reverseDigits(list []int) []int {
	reverse := make([]int, len(list))
	for i, v := range list {
		reverse[len(list)-i-1] = v
	}
	return reverse
}

func splitDigits(x int) []int {
	digits := []int{}
	if x == 10 {
		return []int{1, 0}
	}
	for x > 0 {
		digit := x % 10
		digits = append(digits, digit)
		x = x / 10
	}
	return digits
}
