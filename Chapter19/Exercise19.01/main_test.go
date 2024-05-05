package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		inputs []int
		want   int
	}{
		{
			name:   "Test Case 1",
			inputs: []int{5, 6},
			want:   11,
		},
		{
			name:   "Test Case 2",
			inputs: []int{11, 7},
			want:   18,
		},
		{
			name:   "Test Case 3",
			inputs: []int{1, 8},
			want:   9,
		},
		{
			name:   "Test Case 4 (intentional failure)",
			inputs: []int{2, 3},
			want:   8, // This should be 5 + 3, intentionally incorrect
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := add(test.inputs[0], test.inputs[1])
			if test.name == "Test Case 4 (intentional failure)" {
				assert.NotEqual(t, test.want, got)
			} else {
				assert.Equal(t, test.want, got)
			}
		})
	}
}
