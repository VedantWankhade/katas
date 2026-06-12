package stack

type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   []int{},
		stack: []int{},
	}
}

func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	if value > this.min[len(this.min)-1] {
		value = this.min[len(this.min)-1]
	}
	this.min = append(this.min, value)
}

func (this *MinStack) Pop() {
	this.min = this.min[:len(this.min)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
