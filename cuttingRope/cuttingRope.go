package main

func main() {
	cuttingRope(10)

}

//尽可能的均等平均分是否是最大的收益呢
//https://leetcode-cn.com/problems/jian-sheng-zi-lcof/

func cuttingRope(n int) int {
	res := 0
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		val := 0
		for j := 2; j < i; j++ {
			temp := max(j*(i-j), dp[i-j]*j)
			if temp > val {
				val = temp
			}
		}
		dp[i] = val
	}
	res = maxArr(dp)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArr(arr []int) int {
	max := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
