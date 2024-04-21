package main

import "testing"

func Test_a(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//TBD
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := a(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("a() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
