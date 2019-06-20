/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.
Example:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
*/
/*
#### 解题思路
- 使用两个栈进行实现，一个in，一个out，所有元素从栈in进入，当元素出去时，先看out栈是否有值，有的话直接pop
- 若out栈为空，从in栈依次将值抛出并压入out栈，再从out栈取值从而达到先入先出的效果
*/

package main

type Stack struct {
	s []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(x int) {
	s.s = append(s.s, x)
}

func (s Stack) Peek() int {
	return s.s[len(s.s)-1]
}

func (s *Stack) Pop() int {
	t := s.Peek()
	s.s = s.s[:len(s.s)-1]
	return t
}

func (s Stack) Size() int {
	return len(s.s)
}

func (s Stack) Empty() bool {
	return len(s.s) == 0
}

type MyQueue struct {
	stackIn  *Stack
	stackOut *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stackIn: NewStack(), stackOut: NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			this.stackOut.Push(this.stackIn.Pop())
		}
	}
	return this.stackOut.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			this.stackOut.Push(this.stackIn.Pop())
		}
	}
	return this.stackOut.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stackOut.Empty() && this.stackIn.Empty()
}
