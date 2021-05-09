package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	maxVal, minVal, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal
		}
		maxVal = max(nums[i]*maxVal, nums[i])
		minVal = min(nums[i]*minVal, nums[i])
		res = max(res, maxVal)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
