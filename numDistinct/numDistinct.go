package main

func main() {

}

func numDistinct(s string, t string) int {
	l1 := len(s)
	l2 := len(t)
	dp := make([][]int, l1+1)
	dp[0] = make([]int, l2+1)
	dp[0][0] = 1
	for i := 1; i <= l2; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = 1
		for j := 1; j <= l2; j++ {
			if s[i-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	return dp[l1][l2]
}
