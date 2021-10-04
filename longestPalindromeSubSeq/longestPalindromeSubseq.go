package main

func main() {

}

//状态优化版本
func longestPalindromeSubseq1(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = temp
		}
	}
	return dp[n-1]
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
