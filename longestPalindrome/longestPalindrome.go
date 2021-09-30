package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindromeDp(s))

}

func longestPalindromeDp(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	maxLen := 0
	maxStart := 0
	maxEnd := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = false
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					if j-i+1 > maxLen {
						maxLen = j - i + 1
						maxStart = i
						maxEnd = j
					}
					dp[i][j] = true
				} else {
					if dp[i+1][j-1] {
						if j-i+1 > maxLen {
							maxLen = j - i + 1
							maxStart = i
							maxEnd = j
						}
						dp[i][j] = true
					}
				}
			}
		}
	}
	//fmt.Println("maxStart:", maxStart, " maxEnd:", maxEnd)
	return s[maxStart : maxEnd+1]
}

func longestPalindrome(s string) string {
	ret := ""
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			subStr := s[i:j]
			if isPalindromic(subStr) {
				if len(subStr) > max {
					max = len(subStr)
					ret = subStr
				}
			}

		}
	}
	return ret

}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
