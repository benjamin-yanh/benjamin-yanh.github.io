package top100liked_golang

import (
	"fmt"
	"sort"
	"testing"
)

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	nums1 = append(nums1, nums2...)
//	sort.Ints(nums1)
//	if len(nums1)%2 == 0 {
//		val1 := float64(nums1[len(nums1)/2])
//		val2 := float64(nums1[len(nums1)/2-1])
//		return (val2 + val1) / 2.0
//	} else {
//		return float64(nums1[len(nums1)/2])
//	}
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		val1 := float64(nums1[len(nums1)/2])
		val2 := float64(nums1[len(nums1)/2-1])
		return (val2 + val1) / 2.0
	} else {
		return float64(nums1[len(nums1)/2])
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	vals := []int{1, 3, 49, 4}
	//fmt.Println(findMedianSortedArrays(vals, []int{2}))
	fmt.Println(vals[:len(vals)-1])
	//sort.Ints(vals[:len(vals)-1])
	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
