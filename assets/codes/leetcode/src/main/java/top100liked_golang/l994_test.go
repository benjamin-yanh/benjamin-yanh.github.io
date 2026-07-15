package top100liked_golang

import (
	"fmt"
	"testing"
)

func orangesRotting(grid [][]int) int {
	res := 0
	rows := len(grid)
	cols := len(grid[0])
	maps := make(map[string]bool)
	queue := [][]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				key := fmt.Sprintf("%v_%v", i, j)
				if _, ok := maps[key]; !ok {
					queue = append(queue, []int{i, j})
				}
				maps[key] = true
			}
		}
	}
	ans := -1
	for len(queue) > 0 {
		length := len(queue)
		ans++
		for m := 0; m < length; m++ {
			ints := queue[m]
			if ints[0]+1 < rows && (grid[ints[0]+1][ints[1]] == 1) {
				key := fmt.Sprintf("%d_%d", ints[0]+1, ints[1])
				if _, ok := maps[key]; !ok {
					queue = append(queue, []int{ints[0] + 1, ints[1]})
				}
				maps[key] = true
			}
			if ints[0]-1 >= 0 && (grid[ints[0]-1][ints[1]] == 1) {
				key := fmt.Sprintf("%d_%d", ints[0]-1, ints[1])
				if _, ok := maps[key]; !ok {
					queue = append(queue, []int{ints[0] - 1, ints[1]})
				}
				maps[key] = true
			}
			if ints[1]+1 < cols && (grid[ints[0]][ints[1]+1] == 1) {
				key := fmt.Sprintf("%d_%d", ints[0], ints[1]+1)
				if _, ok := maps[key]; !ok {
					queue = append(queue, []int{ints[0], ints[1] + 1})
				}
				maps[key] = true
			}
			if ints[1]-1 >= 0 && (grid[ints[0]][ints[1]-1] == 1) {
				key := fmt.Sprintf("%d_%d", ints[0], ints[1]-1)
				if _, ok := maps[key]; !ok {
					queue = append(queue, []int{ints[0], ints[1] - 1})
					maps[key] = true
				}
			}
			grid[ints[0]][ints[1]] = 2
		}
		queue = queue[length:]
	}
	if ans > 0 {
		res += ans
	}

	for _, ints := range grid {
		for _, ele := range ints {
			if ele == 1 {
				return -1
			}
		}
	}

	return res
}

func Test_orangesRotting(t *testing.T) {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	orangesRotting(grid)
}
