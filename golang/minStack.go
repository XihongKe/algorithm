package golang

import "fmt"

type MinStack struct {
	stack       []int
	helperStack []int
}

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/
/*
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.helperStack) == 0 || this.helperStack[len(this.helperStack)-1] >= x {
		this.helperStack = append(this.helperStack, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.stack[len(this.stack)-1] == this.helperStack[len(this.helperStack)-1] {
		this.helperStack = this.helperStack[0 : len(this.helperStack)-1]
	}
	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	fmt.Printf("stack=%#v \r\nhelerStack=%#v \r\n", this.stack, this.helperStack)
	return this.helperStack[len(this.helperStack)-1]
}
