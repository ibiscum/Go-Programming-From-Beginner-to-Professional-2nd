package main

import "testing"

func Test_typeExample(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeExample(tt.args.i)
		})
	}
}
