package payroll

import "testing"

func TestPayDetails(t *testing.T) {
	type args struct {
		p Payer
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PayDetails(tt.args.p)
		})
	}
}
