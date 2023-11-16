package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	for i := 0; i < x; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return x

}

func main() {
}
