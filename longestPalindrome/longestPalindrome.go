package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))

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
