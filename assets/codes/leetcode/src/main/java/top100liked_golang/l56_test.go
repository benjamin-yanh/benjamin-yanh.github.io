package top100liked_golang

import (
	"fmt"
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})

	result := make([][]int, 0)
	result = append(result, intervals[0])
	for _, interval := range intervals {
		tail := result[len(result)-1]
		if (tail[0] <= interval[0] && interval[0] <= tail[1]) ||
			(tail[1] <= interval[1] && interval[1] <= tail[1]) ||
			(interval[0] <= tail[0] && tail[0] >= interval[1]) ||
			(interval[0] <= tail[1] && tail[1] >= interval[1]) {
			result[len(result)-1][0] = min(tail[0], tail[1], interval[0], interval[1])
			result[len(result)-1][1] = max(tail[0], tail[1], interval[0], interval[1])
		} else {
			result = append(result, interval)
		}
	}

	return result
}

func Test_merge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))
}
