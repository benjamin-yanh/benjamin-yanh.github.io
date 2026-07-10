package top100liked_golang

import (
	"fmt"
	"testing"
)

//
//func searchMatrix(matrix [][]int, target int) bool {
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[0]); j++ {
//			if matrix[i][j] == target {
//				return true
//			}
//		}
//	}
//	return false
//}

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	x, y := 0, cols-1

	for x < rows && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}

	return false
}

func Test_searchMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 8))
	fmt.Println(searchMatrix([][]int{{1}}, 8))
}
