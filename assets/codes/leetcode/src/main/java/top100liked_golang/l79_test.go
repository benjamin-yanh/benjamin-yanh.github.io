package top100liked_golang

import (
	"fmt"
	"testing"
)

func exist(board [][]byte, word string) bool {
	var char = word[0]
	starts := [][]int{}
	for i, bytes := range board {
		for j, b := range bytes {
			if b == char {
				starts = append(starts, []int{i, j})
			}
		}
	}
	if len(starts) == 0 {
		return false
	}

	rows := len(board)
	cols := len(board[0])
	for _, start := range starts {
		back := board[start[0]][start[1]]
		board[start[0]][start[1]] = '0'
		val := findExist(board, start[0], start[1], rows, cols, word[1:])
		board[start[0]][start[1]] = back
		if val {
			return true
		}
	}

	return false

}

func findExist(board [][]byte, row, col, rows, cols int, word string) bool {
	if len(word) == 0 {
		return true
	}

	if row-1 >= 0 && board[row-1][col] == word[0] {
		back := board[row-1][col]
		board[row-1][col] = '0'
		if findExist(board, row-1, col, rows, cols, word[1:]) {
			return true
		}
		board[row-1][col] = back
	}

	if row+1 < rows && board[row+1][col] == word[0] {
		back := board[row+1][col]
		board[row+1][col] = '0'
		if findExist(board, row+1, col, rows, cols, word[1:]) {
			return true
		}
		board[row+1][col] = back
	}
	if col-1 >= 0 && board[row][col-1] == word[0] {
		back := board[row][col-1]
		board[row][col-1] = '0'
		if findExist(board, row, col-1, rows, cols, word[1:]) {
			return true
		}
		board[row][col-1] = back
	}
	if col+1 < cols && board[row][col+1] == word[0] {
		back := board[row][col+1]
		board[row][col+1] = '0'
		if findExist(board, row, col+1, rows, cols, word[1:]) {
			return true
		}
		board[row][col+1] = back
	}

	return false
}

func Test_exist(t *testing.T) {
	//str := "SEE"
	//val := exist([][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//}, str)
	//fmt.Println(val)
	str := "ba"
	val := exist([][]byte{
		{'a', 'b'},
	}, str)
	fmt.Println(val)
}
