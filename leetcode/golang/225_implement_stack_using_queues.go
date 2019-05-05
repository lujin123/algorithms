package leetcode

// https://leetcode-cn.com/problems/implement-stack-using-queues/
/**
// 225. 用队列实现栈

使用队列实现栈的下列操作：
	push(x) -- 元素 x 入栈
	pop() -- 移除栈顶元素
	top() -- 获取栈顶元素
	empty() -- 返回栈是否为空

注意:
1. 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
2. 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
3. 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
*/
/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type MyStack struct {
	elements []int
}

/** Initialize your data structure here. */
func NewMyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.elements = append(this.elements, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.elements)
	e := this.elements[n-1]
	this.elements = this.elements[:n-1]
	return e
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.elements[len(this.elements)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.elements) == 0
}
