package main

import (
	"fmt"
	"math"
)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	for i := 0; ; i++ {
		v := math.Pow(2, float64(i))
		if v >= float64(n) {
			return i
		}
	}
}

func main() {
	fmt.Println(climbStairs(4))
}
