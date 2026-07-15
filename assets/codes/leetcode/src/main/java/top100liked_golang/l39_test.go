package top100liked_golang

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	maps := make(map[string]bool)
	findCombinationSum(candidates, []int{}, target, &ans, maps)
	return ans
}
func JoinInts(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ",")
}

func findCombinationSum(candidates []int, result []int, target int, ans *[][]int, exist map[string]bool) {
	if target == 0 {
		value := append([]int{}, result...)
		sort.Ints(value)
		key := JoinInts(value)
		if _, ok := exist[key]; !ok {
			*ans = append(*ans, value)
			exist[key] = true
		}
		return
	}
	if target < 0 {
		return
	}
	length := len(candidates)
	for i := 0; i < length; i++ {
		result = append(result, candidates[i])
		findCombinationSum(candidates, result, target-candidates[i], ans, exist)
		result = result[:len(result)-1]
	}
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
