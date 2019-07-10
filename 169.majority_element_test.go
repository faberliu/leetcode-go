package main

import (
	"testing"
)

func Test_majorityElement_1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{3, 2, 3}}, 3},
		{"test2", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
		{"test2", args{[]int{3, 3, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement1(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{3, 2, 3}}, 3},
		{"test2", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
		{"test2", args{[]int{3, 3, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement4(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{3, 2, 3}}, 3},
		{"test2", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
		{"test2", args{[]int{3, 3, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement4(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{3, 2, 3}}, 3},
		{"test2", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
		{"test2", args{[]int{3, 3, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement3(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement3() = %v, want %v", got, tt.want)
			}
		})
	}
}
