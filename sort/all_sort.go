package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 4}
	//quickSort(nums, 0, 4)
	createHeap(nums)
	fmt.Println(nums)

}

func quickSort(nums []int, l, r int) {
	if l < r {
		mid := partion(nums, l, r)
		quickSort(nums, l, mid)
		quickSort(nums, mid+1, r)
	}
}

func partion(nums []int, l, r int) int {
	temp := nums[l]

	for l < r {
		for l < r && nums[r] >= temp {
			r--
		}
		fmt.Println("r:", r)
		nums[l] = nums[r]
		for l < r && nums[l] <= temp {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = temp
	return l
}

func partionTopK(nums []int, l, r int) int {
	temp := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] > temp {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1

}

func heapSort(nums int) {

}

func createHeap(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		minHeap(nums, i, n)
	}
	fmt.Println(nums)

}

func minHeap(nums []int, posIndex, n int) {
	minIndex := posIndex
	leftIndex := posIndex*2 + 1
	rightIndex := posIndex*2 + 2
	if leftIndex < n && nums[leftIndex] < nums[minIndex] {
		minIndex = leftIndex
	}
	if rightIndex < n && nums[rightIndex] < nums[minIndex] {
		minIndex = rightIndex
	}
	if minIndex != posIndex {
		nums[posIndex], nums[minIndex] = nums[minIndex], nums[posIndex]
		minHeap(nums, minIndex, n)
	}

}
