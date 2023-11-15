package main

import "testing"

type args struct {
	nums   []int
	target int
}

type testpair struct {
	value  args
	answer int
}

var test = []testpair{
	testpair{value: args{nums: []int{1, 3, 5, 6}, target: 5}, answer: 2},
	testpair{value: args{nums: []int{1, 3, 5, 6}, target: 2}, answer: 1},
	testpair{value: args{nums: []int{1, 3, 5, 6}, target: 7}, answer: 4},
	testpair{value: args{nums: []int{2, 5}, target: 1}, answer: 0},
	testpair{value: args{nums: []int{-1, 3, 5, 6}, target: 0}, answer: 1},
}

func TestSearchInsert(t *testing.T) {
	for _, pair := range test {
		v := searchInsert(pair.value.nums, pair.value.target)
		if v != pair.answer {
			t.Error(
				"For", pair.value.nums, pair.value.target,
				"expected", pair.answer,
				"got", v,
			)
		}
	}

}
