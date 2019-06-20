/*
Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.
Example:

MyStack stack = new MyStack();

stack.push(1);
stack.push(2);
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false
Notes:

You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).
*/

/*
# 解题思路
1. 用两个队列in out进行存储
2. push时永远push到in
3. 当pop和top时，先从队列in中将值复制到队列out中，并将最后一个值输出(移除)；若队列in为空，从out中pop数据到队列in中，最后一个值输出(删除)
*/

package main

type Queue struct {
	s []int
}

func (m *Queue) Push(x int) {
	m.s = append(m.s, x)
}

func (m *Queue) Pop() int {
	temp := m.s[0]
	m.s = m.s[1:]
	return temp
}

func (m *Queue) Peek() int {
	return m.s[0]
}

func (m *Queue) Size() int {
	return len(m.s)
}

func (m *Queue) Empty() bool {
	return len(m.s) == 0
}

type MyStack struct {
	queueIn  Queue
	queueOut Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queueIn.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if !this.queueIn.Empty() {
		for this.queueIn.Size() > 1 {
			this.queueOut.Push(this.queueIn.Pop())
		}
		return this.queueIn.Pop()
	}
	for this.queueOut.Size() > 1 {
		this.queueIn.Push(this.queueOut.Pop())
	}
	return this.queueOut.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if !this.queueIn.Empty() {
		for this.queueIn.Size() > 1 {
			this.queueOut.Push(this.queueIn.Pop())
		}
		return this.queueIn.Peek()
	}
	for this.queueOut.Size() > 1 {
		this.queueIn.Push(this.queueOut.Pop())
	}
	pop := this.queueOut.Pop()
	this.queueIn.Push(pop)
	return pop
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queueIn.Empty() && this.queueOut.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
