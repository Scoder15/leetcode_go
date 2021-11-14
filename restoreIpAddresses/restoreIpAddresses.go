package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))

}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	track := make([]string, 0)
	backTrack(&res, track, s, 0)
	return res

}

func backTrack(res *[]string, track []string, s string, start int) {
	fmt.Println("track:", track)
	if start == len(s) && len(track) == 4 {
		temp := track[0] + "." + track[1] + "." + track[2] + "." + track[3]
		*res = append(*res, temp)
	}

	for i := start; i < len(s); i++ {
		track = append(track, s[start:i+1])
		if i-start+1 <= 3 && len(track) <= 4 && isIpAddress(s, start, i) {
			backTrack(res, track, s, i+1)
		} else {
			return
		}
		track = track[:len(track)-1]
	}

}

func isIpAddress(s string, start, end int) bool {
	sNum, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if sNum < 0 || sNum > 255 {
		return false
	}
	return true
}
