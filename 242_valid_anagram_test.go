package main

import (
	"testing"
)

func Test_isAnagram1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test_isanagram1", args{"abdde", "bdeda"}, true},
		{"test_isanagram1", args{"acdeff", "acdef"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test_isanagram2", args{"abdde", "bdeda"}, true},
		{"test_isanagram2", args{"acdeff", "acdef"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram2() = %v, want %v", got, tt.want)
			}
		})
	}
}
