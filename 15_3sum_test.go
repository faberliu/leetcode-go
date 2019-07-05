/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package main

import (
	"reflect"
	"testing"
)

func Test_threeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"a", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{"b", args{[]int{0, 0, 0, 0}}, [][]int{{0, 0, 0}}},
		{"c", args{[]int{0, 2, 2, 3, 0, 1, 2, 3, -1, -4, 2}}, [][]int{{-4, 1, 3}, {-4, 2, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"a", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{"b", args{[]int{0, 0, 0, 0}}, [][]int{{0, 0, 0}}},
		{"c", args{[]int{0, 2, 2, 3, 0, 1, 2, 3, -1, -4, 2}}, [][]int{{-4, 1, 3}, {-4, 2, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
