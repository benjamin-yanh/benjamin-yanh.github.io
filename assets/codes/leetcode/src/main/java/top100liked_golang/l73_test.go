package top100liked_golang

import (
	"slices"
	"testing"
)

func setZeroes(matrix [][]int) {
	col := []int{}
	row := []int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				col = append(col, j)
				row = append(row, i)
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if slices.Contains(col, j) || slices.Contains(row, i) {
				matrix[i][j] = 0
			}
		}
	}

}

func Test_setZeroes(t *testing.T) {

}
