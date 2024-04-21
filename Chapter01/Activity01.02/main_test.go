package main

import "testing"

func Test_swap(t *testing.T) {
	type args struct {
		a *int
		b *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.a, tt.args.b)
		})
	}
}
