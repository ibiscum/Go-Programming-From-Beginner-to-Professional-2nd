package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		name      string
		input     int
		wantedInt int
		wantedStr string
	}{

		{
			name:      "Even",
			input:     6,
			wantedInt: 6,
			wantedStr: "Even",
		},

		{
			name:      "Even",
			input:     10,
			wantedInt: 10,
			wantedStr: "Even",
		},

		{
			name:      "Odd",
			input:     15,
			wantedInt: 15,
			wantedStr: "Odd",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotInt, gotStr := checkNumbers(tc.input)

			if gotInt != tc.wantedInt || gotStr != tc.wantedStr {
				t.Errorf("wanted integer: %v, got integer: %v, wanted string %v, got string: %v ", tc.wantedInt, gotInt, tc.wantedStr, gotStr)
			}

		})

	}

}
