package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permute(nums))

}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	visited := make(map[int]bool, 0)
	sort.Ints(nums)
	backtrack(&res, nums, track, &visited)
	return res
}

func backtrack(res *[][]int, nums []int, track []int, visited *map[int]bool) {
	if len(track) == len(nums) {
		fmt.Println(track)
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && (*visited)[nums[i-1]] == false {
			continue
		}
		if (*visited)[nums[i]] {
			continue
		}
		(*visited)[nums[i]] = true
		track = append(track, nums[i])
		backtrack(res, nums, track, visited)
		track = track[:len(track)-1]
		(*visited)[nums[i]] = false
	}
}
