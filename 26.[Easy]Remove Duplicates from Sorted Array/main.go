package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	hashMap := map[int]bool{}
	index := 0
	for _, v := range nums {
		// fmt.Println(hashMap[v], v)
		if hashMap[v] == false {
			hashMap[v] = true
			nums[index] = v
			index++
		}
	}
	return len(hashMap)
}
