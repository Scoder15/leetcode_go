package main

func main() {

}

func lastStoneWeightII(stones []int) int {
	dp := make([]int, 15001)
	for i := 0; i < 15001; i++ {
		dp[i] = 0
	}
	n := len(stones)
	sum := 0
	for i := 0; i < n; i++ {
		sum += stones[i]
	}
	target := sum / 2
	for i := 0; i < n; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target]*2

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
