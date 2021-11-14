package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partion(s))

}

func partion(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res := make([][]string, 0)
	track := make([]string, 0)
	backTrack(&res, track, s, 0)
	return res
}

func backTrack(res *[][]string, track []string, s string, startIndex int) {
	if startIndex >= len(s) {
		temp := make([]string, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}
	for i := startIndex; i < len(s); i++ {
		str := s[startIndex : i+1]
		if isPalindromicStr(str) {
			track = append(track, str)
			backTrack(res, track, s, i+1)
			track = track[:len(track)-1]
		} else {
			continue
		}
	}
}

func isPalindromicStr(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
