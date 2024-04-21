package main

import (
	"testing"
)

func Test_cat_Speak(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cat{
				name: tt.fields.name,
			}
			if got := c.Speak(); got != tt.want {
				t.Errorf("cat.Speak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emptyDetails(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyDetails(tt.args.i)
		})
	}
}

func Test_catDetails(t *testing.T) {
	type args struct {
		i Speaker
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			catDetails(tt.args.i)
		})
	}
}
