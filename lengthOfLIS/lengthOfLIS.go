package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
	fmt.Println(findLengthOfLCIS(nums))
}

//非连续上升子序列
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, len(nums)+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)

			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	dp := make([]int, len(nums)+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		if dp[i+1] > res {
			res = dp[i+1]
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
