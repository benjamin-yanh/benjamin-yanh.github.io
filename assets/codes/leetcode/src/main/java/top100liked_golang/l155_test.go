package top100liked_golang

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{}}
}

func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, value)
	} else {
		this.minStack = append(this.minStack, min(value, this.minStack[len(this.minStack)-1]))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
