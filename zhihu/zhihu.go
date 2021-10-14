package main

func main() {

}

var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func isLandNums(nums [][]byte) int {
	//
	m := len(nums)
	n := len(nums)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, 0)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums[i][j] == '1' && !visited[i][j] {
				dfs(nums, &visited, i, j)
				res++
			}
		}
	}
	return res

}

func dfs(nums [][]byte, visited *[][]bool, x, y int) {
	(*visited)[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(nums, nx, ny) && !(*visited)[nx][ny] && nums[nx][ny] == '1' {
			dfs(nums, visited, nx, ny)
		}
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) {
		return true
	}
	return false
}
