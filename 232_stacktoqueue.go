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
