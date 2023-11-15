package main

import "testing"

type testpair struct {
	value  string
	answer int
}

var test = []testpair{
	// testpair{value: "Hello World", answer: 5},
	testpair{value: "   fly me   to   the moon  ", answer: 4},
	// testpair{value: "luffy is still joyboy", answer: 6},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, pair := range test {
		v := lengthOfLastWord(pair.value)
		if v != pair.answer {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
		}
	}

}
