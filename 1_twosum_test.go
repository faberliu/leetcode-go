package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	index := twoSum([]int{2, 7, 11, 15}, 9)
	result := []int{0, 1}
	if !reflect.DeepEqual(index, []int{0, 1}) {
		t.Error("testing failed.result is", index, " and expect", result)
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"正常测试", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"异常测试", args{[]int{2, 7, 11, 15}, 10}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
