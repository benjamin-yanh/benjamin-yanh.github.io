package top100liked_golang

import (
	"fmt"
	"testing"
)

func spiralOrder(matrix [][]int) []int {
	directors := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := []int{}
	length := len(matrix) * len(matrix[0])

	row := 0
	col := 0
	director := 0
	for i := 0; i < length; i++ {
		result = append(result, matrix[row][col])
		matrix[row][col] = -101

		nextRow := row + directors[director][0]
		nextCol := col + directors[director][1]

		if nextRow >= len(matrix) || nextRow < 0 || nextCol >= len(matrix[0]) || nextCol < 0 ||
			matrix[nextRow][nextCol] < -100 {
			director = (director + 1) % 4
		}

		row += directors[director][0]
		col += directors[director][1]

	}

	return result

}

func Test_spiralOrder(t *testing.T) {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
