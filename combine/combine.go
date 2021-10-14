package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3}
	// aa := make([]int, 3)
	// copy(aa, nums)
	// fmt.Println("aa:", aa, " nums:", nums)
	fmt.Println(combine(4, 2))
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println("slice1:", slice1, " slice2:", slice2)

}

func combine(n, k int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	backTrack(track, &result, 1, n, k)
	return result
}

func backTrack(track []int, result *[][]int, start, n, k int) {
	if len(track) == k {
		temp := make([]int, len(track))
		copy(temp, track)
		// fmt.Println("temp:", temp, " track:", track)
		//temp = track
		(*result) = append(*result, temp)
		// 	fmt.Println(" result:", *result)
		// 	(*result) = append(*result, track)
		// 	fmt.Println(" track:", track, " result:", *result)
	}
	i := start
	for i <= n {
		track = append(track, i)
		backTrack(track, result, i+1, n, k)
		track = track[:len(track)-1]
		i++
	}
	return
}
