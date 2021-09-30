package main

func main() {

}

func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}
	res := 0
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
