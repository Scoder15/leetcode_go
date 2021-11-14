package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))

}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	track := make([]int, 0)
	sort.Ints(candidates)
	backTrack1(&res, track, candidates, target, 0)
	return res

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
		if i > startIndex && candidates[i-1] == candidates[i] {
			continue
		}
		if target >= candidates[i] {
			track = append(track, candidates[i])
			backTrack1(res, track, candidates, target-candidates[i], i+1)
			track = track[:len(track)-1]
		}
	}

}
