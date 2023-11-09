package main

import "testing"

type testpair struct {
	value  []string
	answer string
}

var testtable = []testpair{
	{value: []string{"flower", "flow", "flight"}, answer: "fl"},
	{value: []string{"dog", "racecar", "car"}, answer: ""},
	{value: []string{"cir", "car"}, answer: "c"},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, pair := range testtable {
		v := longestCommonPrefix(pair.value)
		if v != pair.answer {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
		}
	}
}
