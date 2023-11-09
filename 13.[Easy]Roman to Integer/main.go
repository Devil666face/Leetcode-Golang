package main

var chars = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	romanToInt("III")
	romanToInt("LVIII")
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	total := 0
	for i, c := range []rune(s) {
		current := chars[string(c)]
		next := 0
		if i+1 < len([]rune(s)) {
			next = chars[string([]rune(s)[i+1])]
		}
		if current < next {
			total -= current
		} else {
			total += current
		}
	}
	return total
}
