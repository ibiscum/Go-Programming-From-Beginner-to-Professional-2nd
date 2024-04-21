package main

import (
	"io"
	"reflect"
	"testing"
)

func Test_loadPerson1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadPerson1(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadPerson1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPerson1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadPerson2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadPerson2(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadPerson2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPerson2() = %v, want %v", got, tt.want)
			}
		})
	}
}
