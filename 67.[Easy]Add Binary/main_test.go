package main

import (
	"testing"
)

type testpair struct {
	value  []string
	answer string
}

var testtable = []testpair{
	{value: []string{"11", "1"}, answer: "100"},
	{value: []string{"1010", "1011"}, answer: "10101"},
}

func TestAddBinary(t *testing.T) {
	for _, pair := range testtable {
		v := addBinary(pair.value[0], pair.value[1])
		if v != pair.answer {
			t.Error(
				"For", pair.value,
				"expected", pair.answer,
				"got", v,
			)
			continue
		}
		t.Log(
			"Success for", pair.value,
			"expected", pair.answer,
			"got", v,
		)
	}
}
