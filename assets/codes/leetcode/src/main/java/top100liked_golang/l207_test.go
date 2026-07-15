package top100liked_golang

import (
	"fmt"
	"testing"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	valid := true
	var dsf func(u int)

	dsf = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dsf(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dsf(i)
		}
	}

	return valid
}

func Test_canFinish(t *testing.T) {
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
