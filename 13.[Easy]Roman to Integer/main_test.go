package main

import "testing"

type testpair struct {
	value  string
	answer int
}

var testtable = []testpair{
	{value: "III", answer: 3},
	{value: "LVIII", answer: 58},
	{value: "MCMXCIV", answer: 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, pair := range testtable {
		v := romanToInt(pair.value)
		if v != pair.answer {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
		}
	}
}
