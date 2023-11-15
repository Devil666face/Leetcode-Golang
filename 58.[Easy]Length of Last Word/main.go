package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	slice := strings.Split(strings.TrimSpace(s), " ")
	return len(slice[len(slice)-1])
}

func main() {
}
