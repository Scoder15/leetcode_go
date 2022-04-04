package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
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
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = '.'
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i, j+1, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = word[k]
	return res
}

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func lastRemaining(n int, m int) int {
	//	return f1(n, m)
	res := 0
	for i := 2; i <= n; i++ {
		res = (m + res) % i
	}
	return res
}

func f1(n, m int) int {
	if n == 1 {
		return 0
	}
	x := f1(n-1, m)
	return (m + x) % n
}

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '+' && i == 0 {
			sign = 1
		} else if v == '-' && i == 0 {
			sign = -1
		} else {
			break
		}
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

	}
	return result * sign

}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || len(nums) == 0 {
		return []int{}
	}
	windows := make([]int, 0, k)
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}
		for len(windows) > 0 && nums[windows[len(windows)-1]] < v {
			windows = windows[0 : len(windows)-1]
		}
		windows = append(windows, i)
		if i >= k-1 {
			result = append(result, nums[windows[0]])
		}
	}
	return result

}

func maxSlidingWindowSum(nums []int, k int) []int {
	windowsSum := 0
	n := len(nums)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		windowsSum += nums[i]
	}
	res = append(res, windowsSum)
	for i := k; i < n; i++ {
		windowsSum += nums[i] - nums[i-k]
		res = append(res, windowsSum)
	}
	return res
}

func minNumber(nums []int) string {
	return ""
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + serialize(root.Left) + "," + serialize(root.Right)
}

func deserialize(list string) *TreeNode {
	ll := strings.Split(list, ",")
	return buildTree(&ll)
}

func buildTree(list *[]string) *TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "#" {
		return nil
	}
	val, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: val}
	root.Left = buildTree(list)
	root.Right = buildTree(list)
	return root
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	stack = append(stack, cur)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			v := stack[0]
			stack = stack[1:]
			res = append(res, v.Val)
			cur = cur.Right
		}
	}
	return res

}

func postorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	stack = append(stack, cur)
	for len(stack) > 0 {
		v := stack[0]
		stack = stack[1:]
		res = append(res, v.Val)
		if v.Left != nil {
			stack = append(stack, v.Left)
		}
		if v.Right != nil {
			stack = append(stack)
		}
	}
	reverseArr(&res)
	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tmp := &ListNode{Val: 0}
	cur := tmp
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return cur.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:len(lists)])
	return mergeTwoLists(left, right)
}

func topKFrequent(nums []int, k int) []int {
	return []int{}

}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	resCount := 0
	left := 0
	right := 0
	start := 0
	end := 0
	need := make(map[byte]int, 0)
	tMap := make(map[byte]bool, 0)
	isInit := true
	for i := 0; i < len(t); i++ {
		if _, ok := need[t[i]]; !ok {
			need[t[i]] = 1
		} else {
			need[t[i]]++
		}
		tMap[t[i]] = true
	}
	//fmt.Println("need:", need)
	needCount := len(t)
	for right < len(s) {
		//subStr := s[left : right+1]
		if _, ok := need[s[right]]; !ok {
			need[s[right]] = -1
		} else {

			if need[s[right]] > 0 {

				needCount--
			}
			need[s[right]]--
		}
		if needCount == 0 {
			for left < right && need[s[left]] < 0 {
				need[s[left]]++
				left++
			}
			if isInit {
				resCount = right - left + 1
				start = left
				end = right + 1
				isInit = false
			}
			if right-left < resCount {
				resCount = right - left + 1
				start = left
				end = right + 1
			}

			if _, ok := need[s[left]]; !ok {
				need[s[left]] = 1
			} else {
				need[s[left]]++
				needCount++
			}
			//fmt.Println("need after:", need, " left:", left, " right:", right)
			left++
		}
		right++
		//fmt.Println("need after:", need, " left:", left, " right:", right, " needCount:", needCount)
	}
	return s[start:end]
}

func isNeed(need map[byte]int, t string) bool {
	for i := 0; i < len(t); i++ {
		if need[t[i]] > 0 {
			return false
		}
	}
	return true
}

func minWindowBruit(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	resCount := len(s) + 1
	res := ""

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subStr := s[i : j+1]
			fmt.Println("subStr:", subStr)
			if isHaveStr(subStr, t) {
				fmt.Println("hhhhh")
				if len(subStr) < resCount {
					res = subStr
					resCount = len(subStr)
				}
			}
		}
	}
	return res
}

func isHaveStr(s string, t string) bool {
	if len(s) < len(t) {
		return false
	}
	m := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = 1
		} else {
			m[s[i]]++
		}
	}
	flag := true
	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; !ok {
			flag = false
		} else {
			m[t[i]]--
			if m[t[i]] < 0 {
				flag = false
			}
		}
	}
	return flag
}

func maximalRectangle(matrix []string) int {
	return 0

}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	dfsMaxSum(root, &maxSum)
	return maxSum
}

func dfsMaxSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := dfsMaxSum(root.Left, maxSum)
	right := dfsMaxSum(root.Right, maxSum)

	innerMaxSum := left + root.Val + right
	*maxSum = max(*maxSum, innerMaxSum)
	outputMaxSum := root.Val + max(left, right)
	return max(outputMaxSum, 0)
}

func countSubstrings(s string, dp *[][]bool) {
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] {
				if j-i < 2 || (*dp)[i+1][j-1] {
					(*dp)[i][j] = true
				}

			}
		}
	}
}

func minCut(s string) int {
	isPalim := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		isPalim[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			isPalim[i][j] = false
		}
	}
	countSubstrings(s, &isPalim)
	//dp初始化
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		dp[i] = i
	}
	for i := 1; i < len(s); i++ {
		if isPalim[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPalim[j+1][i] {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}

	return dp[len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	res := 0
	cache := make([][]int, m)
	//初始化为0
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//visited数组需要重新置零
			//if visited[i][j] == 0 {
			dfs112(i, j, matrix, cache, math.MinInt64)
			res = max(res, cache[i][j])
		}
	}
	return res
}

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs112(x, y int, matrix [][]int, cache [][]int, lastNum int) int {
	//(*visited)[x][y] = 1
	if matrix[x][y] <= lastNum {
		return 0
	}
	if cache[x][y] > 0 {
		return cache[x][y]
	}
	count := 1
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := x + dir[i][1]
		if isBoard(nx, ny, len(matrix), len(matrix[0])) {
			count = max(count, dfs112(nx, ny, matrix, cache, matrix[x][y])+1)
		}
	}
	cache[x][y] = count
	return count
}

func isBoard(x, y int, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	track := make([]int, 0)
	sum := target
	dfs34(root, &track, &res, sum)
	return res
}

func dfs34(root *TreeNode, track *[]int, res *[][]int, sum int) {
	if root == nil {
		return
	}
	*track = append(*track, root.Val)
	if sum == root.Val && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
	}

	dfs34(root.Left, track, res, sum-root.Val)
	dfs34(root.Right, track, res, sum-root.Val)

	*track = (*track)[1 : len(*track)-1]
}

func reverseSwitchStr(s string) string {
	res := ""
	return res
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
	//	fmt.Println(reverseWords("a good   example"))
	//fmt.Println(strToInt("-91283472332"))
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// fmt.Println(maxSlidingWindow(nums, 3))
	// s := "ADOBECODEBANC"
	// t := "ABC"
	// s := "a"
	// t := "a"
	s := "cabwefgewcwaefgcf"
	t := "cae"
	fmt.Println(minWindow(s, t))
}
