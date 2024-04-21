package example1002

import (
	"reflect"
	"testing"
)

func Test_explode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := explode(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("explode() = %v, want %v", got, tt.want)
			}
		})
	}
}
