package main

import (
	"fmt"
	"sort"
)

//"a-z = 97,122"
//"0-9 = 48,57"
func main() {
	//	fmt.Println(letterCombinations("23"))
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))

}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	track := make([]int, 0)
	sort.Ints(candidates)
	backTrack1(&res, track, candidates, target, 0)
	return res
	// if len(candidates) == 0 {
	// 	return [][]int{}
	// }
	// c, res := []int{}, [][]int{}
	// sort.Ints(candidates)
	// findcombinationSum(candidates, target, 0, c, &res)
	// return res

}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target { // 这里可以剪枝优化
			break
		}
		c = append(c, nums[i])
		findcombinationSum(nums, target-nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}

func backTrack1(res *[][]int, track []int, candidates []int, target, startIndex int) {
	if target <= 0 {
		if target == 0 {
			temp := make([]int, len(track))
			copy(temp, track)
			(*res) = append(*res, temp)

		}

		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		track = append(track, candidates[i])
		backTrack1(res, track, candidates, target-candidates[i], i)
		track = track[:len(track)-1]
	}

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
