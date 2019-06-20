/*
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
*/

/*
#### 解题思路
- 维护一个k大小的小顶堆，每次返回堆顶即可
- golang没有现成的小顶堆，通过实现heap.Interface来实现一个int型的小顶堆
*/

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Peek() int {
	return h[0]
}

type KthLargest struct {
	h IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{k: k, h: nums}
	heap.Init(&h.h)
	for h.h.Len() > k {
		heap.Pop(&h.h)
	}
	return h
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(&this.h, val)
	} else {
		if val > this.h.Peek() {
			heap.Pop(&this.h)
			heap.Push(&this.h, val)
		}
	}
	return this.h.Peek()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))  // returns 4
	fmt.Println(obj.Add(5))  // returns 5
	fmt.Println(obj.Add(10)) // returns 5
	fmt.Println(obj.Add(9))  // returns 8
	fmt.Println(obj.Add(4))  // returns 8
}
