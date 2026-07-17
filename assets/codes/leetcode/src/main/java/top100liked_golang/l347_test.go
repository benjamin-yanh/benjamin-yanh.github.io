package top100liked_golang

import "sort"

func topKFrequent(nums []int, k int) []int {
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num] = maps[num] + 1
	}
	values := []int{}
	for _, val := range maps {
		values = append(values, val)
	}
	sort.Ints(values)
	value := values[len(values)-k]
	ans := []int{}
	for key, val := range maps {
		if val >= value {
			ans = append(ans, key)
		}
	}
	return ans
}
