package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println(zeroOne(weight, value, 4))

}

func zeroOne(weight, value []int, bagWeight int) int {
	n := len(weight)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, bagWeight+1)
		for j := 0; j <= bagWeight; j++ {
			if i == 0 && j >= weight[i] {
				dp[i][j] = value[0]
			} else {
				dp[i][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])

			}
		}

	}
	return dp[n-1][bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
