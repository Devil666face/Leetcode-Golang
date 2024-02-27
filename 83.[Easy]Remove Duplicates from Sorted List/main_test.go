package main

import (
	"testing"
)

type testpair struct {
	value  int
	answer int
}

var testtable = []testpair{
	{value: 4, answer: 2},
	{value: 8, answer: 2},
	{value: 36, answer: 6},
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMySqrt(t *testing.T) {
	for _, pair := range testtable {
		v := mySqrt(pair.value)
		if v != pair.answer {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
		}
	}
}
