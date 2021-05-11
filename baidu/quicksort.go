package main

func main() {

}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	k := partion(nums, 0, len(nums)-1)
	quickSort(nums, left, k-1)
	quickSort(nums, k+1, right)
}

func partion(nums []int, left, right int) int {
	privot := nums[left]
	j := left
	i := j - 1
	for j = left; j < right; i++ {
		if nums[j] >= privot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
