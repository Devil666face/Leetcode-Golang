package main

import (
	"testing"
)

type testpair struct {
	value  []int
	answer []int
}

var testtable = []testpair{
	{value: []int{1, 2, 3}, answer: []int{1, 2, 4}},
	{value: []int{4, 3, 2, 1}, answer: []int{4, 3, 2, 2}},
	{value: []int{9}, answer: []int{1, 0}},
	{value: []int{9, 9}, answer: []int{1, 0, 0}},
	{value: []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3}, answer: []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 4}},
	{value: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}, answer: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
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

func TestPlusOne(t *testing.T) {
	for _, pair := range testtable {
		v := plusOne(pair.value)
		if !compare(v, pair.answer) {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
		}
	}
}
