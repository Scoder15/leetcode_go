package main

func main() {

}

func rotate(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	column := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := i + 1; j < column; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	half := (column) / 2
	for i := 0; i < row; i++ {
		for j := 0; j < half; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][column-j-1]
			matrix[i][column-j-1] = tmp
		}
	}

}
