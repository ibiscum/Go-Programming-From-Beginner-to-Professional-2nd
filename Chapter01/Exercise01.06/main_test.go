package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_getConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  bool
		want1 string
		want2 time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := getConfig()
			if got != tt.want {
				t.Errorf("getConfig() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getConfig() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("getConfig() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
