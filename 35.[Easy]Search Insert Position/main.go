package main

func main() {
	searchInsert([]int{1, 3, 5, 6}, 5) //2
	searchInsert([]int{1, 3, 5, 6}, 2) //1
	searchInsert([]int{1, 3, 5, 6}, 7) //4
}

func searchInsert(nums []int, target int) int {
	wasIter := false
	expect := 0
	for i, v := range nums {
		if v == target {
			return i
		}
		if v < target {
			wasIter = true
			expect = i
		}
	}
	if wasIter {
		return expect + 1
	}
	return expect
}
