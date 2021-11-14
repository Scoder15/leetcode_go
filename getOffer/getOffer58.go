package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// s := "lrloseumgh"
	// k := 6
	// fmt.Println(reverseLeftWords(s, k))
	// nums := []int{2, 3, 1, 0, 2, 5, 3}
	// fmt.Println(findRepeartNumber(nums))
	//nums := [][]int{{1, 2, 5}, {3, 2, 1}}
	//fmt.Println(maxValue(nums))
	//num := 12258
	//fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
	fmt.Println(reverseWords("a good   example"))
}

func reverseLeftWords(s string, n int) string {
	s1 := s[n:]
	s2 := s[:n]
	s1 = s1 + s2
	return s1
}

func missnumber(nums []int) int {
	i := 0
	j := len(nums)
	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

func firstUniqChar(s string) byte {
	m := make(map[byte]int, 0)
	ss := []byte(s)
	for _, v := range ss {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	for _, v := range ss {
		if m[v] == 1 {
			return v
		}
	}
	return ' '
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
}

func levelOrderII(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		track := make([]int, 0)
		nextRes := make([]*TreeNode, 0)
		for _, v := range queue {
			track = append(track, v.Val)
			if v.Left != nil {
				nextRes = append(nextRes, v.Left)
			}

			if v.Right != nil {
				nextRes = append(nextRes, v.Right)
			}

		}
		l := len(queue)
		queue = append(queue, nextRes...)
		queue = queue[l:]
		res = append(res, track)
	}
	return res
}

func levelOrderIII(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	count := 0
	for len(queue) > 0 {
		track := make([]int, 0)
		nextRes := make([]*TreeNode, 0)
		for _, v := range queue {
			track = append(track, v.Val)
			if v.Left != nil {
				nextRes = append(nextRes, v.Left)
			}

			if v.Right != nil {
				nextRes = append(nextRes, v.Right)
			}

		}
		l := len(queue)
		queue = append(queue, nextRes...)
		count++
		if count%2 == 0 {
			reverseArr(&track)

		}
		queue = queue[l:]
		res = append(res, track)
	}
	return res
}

func reverseArr(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-i-1] = (*arr)[n-i-1], (*arr)[i]
	}

}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return recur(p.Left, q.Right) && recur(p.Right, q.Left)

}

func fib(n int) int {
	//dp[i] = dp[i-1] + dp[i-2]
	f1 := 1
	f2 := 1
	res := -1
	for i := 3; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}

func maxSubArray(nums []int) int {
	cur := nums[0]
	pre := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = nums[i]
		if pre > 0 {
			cur += pre
		}
		if cur > result {
			result = cur
		}
		pre = cur
	}
	return result
}

func maxSubArray1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	result := dp[0]
	for _, v := range dp {
		if v > result {
			result = v
		}

	}
	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	fmt.Println("m:", m, " n:", n)
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			fmt.Println("i:", i, " j:", j)
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		subStr := s[i-2 : i]
		dp[i] = dp[i-1]
		if isTranslate(subStr) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func isTranslate(s string) bool {
	v, _ := strconv.Atoi(s)
	if v >= 10 && v <= 25 {
		return true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	pre := 1
	cur := 1
	res := 1
	hMap := make(map[byte]int, 0)
	hMap[s[0]] = 0
	i := -1
	for j := 1; j < n; j++ {
		if _, ok := hMap[s[j]]; !ok {
			cur = pre + 1
		} else {
			i = hMap[s[j]]
			if j-i > pre {
				cur = pre + 1
			} else {
				cur = j - i
			}

		}
		hMap[s[j]] = j
		pre = cur
		if cur > res {
			res = cur
		}

	}
	return res
}

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//	res := make([]int, 0)
	n := len(nums)
	i := 0
	j := n - 1
	//fmt.Println("i:", i, " j:", j)
	for i < j {
		for j >= 0 && nums[j]%2 == 0 {
			j--
		}
		for i < n && nums[i]%2 == 1 {
			i++
		}
		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
		//fmt.Println("i:", i, " j:", j)
	}
	return nums
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	l := 0
	r := n - 1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{nums[l], nums[r]}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return []int{}

}

func search(nums []int, l, r, target int) int {
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return nums[mid]
		}
	}
	return -1
}

func reverseWords(s string) string {
	ss := strings.TrimSpace(s)

	sss := strings.Split(ss, " ")
	//reverseStr(&sss)
	fmt.Println("sss", len(sss))
	n := len(sss)
	j := n - 1
	res := ""
	for j >= 0 {
		if sss[j] == " " {
			j--
			continue

		}
		if j == n-1 {
			res = res + sss[j]
			j--
			continue
		}
		res = res + " " + sss[j]
		j--
	}
	return res

}

func reverseStr(s *[]string) {
	n := len(*s)
	for i := 0; i < n/2; i++ {
		(*s)[i], (*s)[n-i-1] = (*s)[n-i-1], (*s)[i]
	}
}

func exist(board [][]byte, word string) bool {
	return false

}
