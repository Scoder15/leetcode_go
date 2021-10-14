package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6}
	//quickSort(nums, 0, 4)
	fmt.Println(findKthLatgest(nums, 4))
	fmt.Println(nums)

}

func findTargetSumWays(nums []int, target int) int {
	ret := 0
	return ret
}

func findKthLatgest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return selectpartion(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, l, r int) {
	if l < r {
		mid := partion(nums, l, r)
		quickSort(nums, l, mid-1)
		quickSort(nums, mid+1, r)
	}
}

func partion(nums []int, l, r int) int {
	//基准元素已经保存勒
	temp := nums[l]
	for l < r {
		for l < r && nums[r] >= temp {
			r--
		}
		//因为基准元素已经保存勒，所以这个可以进行覆盖掉
		nums[l] = nums[r]
		for l < r && nums[l] <= temp {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = temp
	return l
}

func selectpartion(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	p := partionheap(nums, l, r)
	if p < k-1 {
		return selectpartion(nums, p+1, r, k)
	} else if p > k-1 {
		return selectpartion(nums, l, p-1, k)
	} else {
		return nums[p]
	}
}

func partionheap(nums []int, l, r int) int {
	//堆排序的topk问题的基准元素是temp
	temp := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] > temp {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[r], nums[i+1] = nums[i+1], nums[r]
	return i + 1

}
