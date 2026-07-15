package top100liked_golang

import (
	"fmt"
	"testing"
)

func numIslands(grid [][]byte) int {
	res := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			queue := [][]int{}
			if grid[i][j] == '1' {
				queue = append(queue, []int{i, j})
				res++
			}
			maps := make(map[string]bool)
			for len(queue) > 0 {
				ints := queue[0]
				queue = queue[1:]
				if ints[0]+1 < rows && grid[ints[0]+1][ints[1]] == '1' {
					key := fmt.Sprintf("%v_%v", ints[0]+1, ints[1])
					if _, ok := maps[key]; !ok {
						queue = append(queue, []int{ints[0] + 1, ints[1]})
					}
					maps[key] = true
				}
				if ints[0]-1 >= 0 && grid[ints[0]-1][ints[1]] == '1' {
					key := fmt.Sprintf("%v_%v", ints[0]-1, ints[1])
					if _, ok := maps[key]; !ok {
						queue = append(queue, []int{ints[0] - 1, ints[1]})
					}
					maps[key] = true
				}
				if ints[1]+1 < cols && grid[ints[0]][ints[1]+1] == '1' {
					key := fmt.Sprintf("%v_%v", ints[0], ints[1]+1)
					if _, ok := maps[key]; !ok {
						queue = append(queue, []int{ints[0], ints[1] + 1})
					}
					maps[key] = true
				}
				if ints[1]-1 >= 0 && grid[ints[0]][ints[1]-1] == '1' {
					key := fmt.Sprintf("%v_%v", ints[0], ints[1]-1)
					if _, ok := maps[key]; !ok {
						queue = append(queue, []int{ints[0], ints[1] - 1})
						maps[key] = true
					}
				}
				grid[ints[0]][ints[1]] = '0'
			}
		}
	}

	return res
}

func Test_numIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
