package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))

}

//非连续的
func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	dp := make([][]int, l1+1)
	dp[0] = make([]int, l2+1)
	for i := 0; i <= l2; i++ {
		dp[0][i] = 0
	}
	for i := 1; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = 0
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
