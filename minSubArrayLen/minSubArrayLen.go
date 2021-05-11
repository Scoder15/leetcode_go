package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right, sum, res := 0, -1, 0, n+1
	for left < n {
		if (right+1) < n && sum < target {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		fmt.Println(sum)
		if sum >= target {
			res = int(math.Min(float64(res), float64(right-left+1)))
		}

	}
	if res == n+1 {
		return 0
	}
	return res
}

//采用滑动窗口
//设置双指针，从前往后滑动·
