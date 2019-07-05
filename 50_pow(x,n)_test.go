package main

import (
	"testing"
)

func Test_myPow_1(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"myPow_1-1", args{2.000, 10}, 1024},
		{"myPow_1-2", args{2.100, 3}, 9.261},
		{"myPow_1-3", args{2.000, -2}, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow_1(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"myPow1", args{2.000, 10}, 1024},
		{"myPow2", args{2.100, 3}, 9.261},
		{"myPow3", args{2.000, -2}, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
