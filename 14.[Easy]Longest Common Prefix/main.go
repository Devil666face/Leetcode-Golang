package main

import "fmt"

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
	longestCommonPrefix([]string{"cir", "car"})
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs {
		sub := substr(prefix, str)
		prefix = sub
	}
	return prefix
}

func substr(x, y string) string {
	fmt.Println(x, y)
	sub := ""
	for i, c := range []rune(x) {
		if i < len([]rune(y)) {
			if c != []rune(y)[i] {
				return sub
			}
			sub += string(c)
		}
	}
	return sub
}
