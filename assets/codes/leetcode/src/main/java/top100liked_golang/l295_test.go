package top100liked_golang

import (
	"container/heap"
	"testing"
)

type MedianFinder struct {
	values IntHeap
}
type IntHeap []int

func Constructor() MedianFinder {
	h := &IntHeap{}
	// 5. 必须手动 Init
	heap.Init(h)
	return MedianFinder{
		values: *h,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.values.Push(num)
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.values)
	if n == 0 {
		return 0.0
	}
	if n%2 == 0 {
		return float64((this.values[n/2] + this.values[n/2-1])) / 2.0
	} else {
		return float64(this.values[n/2])
	}
}

func (this IntHeap) Len() int {
	return len(this)
}
func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}
func (this *IntHeap) Pop() any {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}
func (this *IntHeap) Push(val interface{}) {
	*this = append(*this, val.(int))
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func Test_MedianFinder(t *testing.T) {
	m := Constructor()
	m.AddNum(3)
}
