package main

import "fmt"

//"a-z = 97,122"
//"0-9 = 48,57"
func main() {
	fmt.Println(letterCombinations("23"))

}

func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	m := make(map[string]string, 0)
	m["2"] = "abc"
	m["3"] = "def"
	m["4"] = "ghi"
	m["5"] = "jkl"
	m["6"] = "mno"
	m["7"] = "pqrs"
	m["8"] = "tuv"
	m["9"] = "wxyz"
	track := make([]rune, 0)
	backTrack2(&ret, track, 0, digits, m)
	return ret
}

func backTrack2(ret *[]string, track []rune, startIndex int, digits string, m map[string]string) {

	if len(track) == len(digits) {
		temp := make([]rune, len(digits))
		copy(temp, track)
		(*ret) = append((*ret), string(temp))
	}
	if startIndex >= len(digits) {
		return
	}

	value := string(digits[startIndex])
	valueStr := m[value]
	for _, i := range valueStr {
		track = append(track, i)
		backTrack2(ret, track, startIndex+1, digits, m)
		track = track[:len(track)-1]
	}

}

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	backTrack(track, &result, 1, n, k)
	return result
}

func backTrack(track []int, result *[][]int, start, n, k int) {
	if n == 0 {
		if len(track) == k {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
		}
		return
	}
	for i := start; i < 10; i++ {
		if n >= i {
			track = append(track, i)
			backTrack(track, result, i+1, n-i, k)
			track = track[:len(track)-1]

		}
	}

}
