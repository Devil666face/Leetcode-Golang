package main

import "testing"

type testpair struct {
	value  int
	answer bool
}

var testtable = []testpair{
	{value: 121, answer: true},
	{value: -121, answer: false},
	{value: 10, answer: false},
}

func TestPalindrome(t *testing.T) {
	for _, pair := range testtable {
		v := isPalindrome(pair.value)
		if v != pair.answer {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
		}
	}
}
