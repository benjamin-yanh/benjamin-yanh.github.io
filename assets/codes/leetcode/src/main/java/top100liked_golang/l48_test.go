package top100liked_golang

import "testing"

func rotate2(matrix [][]int) {
	rows := len(matrix)

	// 水平翻转
	for i := 0; i < rows/2; i++ {
		matrix[i], matrix[rows-1-i] = matrix[rows-1-i], matrix[i]
	}
	// 对角线翻转
	for i := 0; i < rows; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return
}

func Test_rotate2(t *testing.T) {
	rotate2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
