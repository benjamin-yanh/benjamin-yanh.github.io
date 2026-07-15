package top100liked_golang

import (
	"strings"
	"testing"
)

func solveNQueens(n int) (ans [][]string) {
	template := make([]string, n)
	for i := range template {
		template[i] = "."
	}
	findSolveNQueens(n, 0, [][]string{}, &ans, template)
	return ans
}

func findSolveNQueens(n, step int, queens [][]string, ans *[][]string, template []string) {
	if n == step {
		if len(queens) < n {
			return
		}
		solve := []string{}
		for _, queen := range queens {
			solve = append(solve, strings.Join(queen, ""))
		}
		*ans = append(*ans, solve)
		return
	}
	for i := 0; i < n; i++ {
		ints := make([]string, n)
		copy(ints, template)
		ints[i] = "Q"
		if findUp(step, i, queens) && findLeftUp(step-1, i-1, queens) && findRightUp(step-1, i+1, queens) {
			queens = append(queens, ints)
			findSolveNQueens(n, step+1, queens, ans, template)
			queens = queens[:len(queens)-1]
		}
	}
}

func findUp(x, y int, queens [][]string) bool {
	for _, queen := range queens {
		if queen[y] == "Q" {
			return false
		}
	}
	return true
}

func findLeftUp(x, y int, queens [][]string) bool {
	for x >= 0 && y >= 0 {
		if queens[x][y] == "Q" {
			return false
		}
		x--
		y--
	}
	return true
}

func findRightUp(x, y int, queens [][]string) bool {
	if len(queens) <= 0 {
		return true
	}
	length := len(queens[x])
	for x >= 0 && y < length {
		if queens[x][y] == "Q" {
			return false
		}
		x--
		y++
	}
	return true
}

func Test_solveNQueens(t *testing.T) {
	solveNQueens(4)
}
