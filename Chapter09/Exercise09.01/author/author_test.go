package author

import (
	"reflect"
	"testing"
)

func TestNewAuthor(t *testing.T) {
	type args struct {
		name    string
		contact string
	}
	tests := []struct {
		name string
		args args
		want *Author
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthor(tt.args.name, tt.args.contact); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthor_WriteChapter(t *testing.T) {
	type fields struct {
		Name    string
		Contact string
	}
	type args struct {
		chapterTitle string
		content      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Author{
				Name:    tt.fields.Name,
				Contact: tt.fields.Contact,
			}
			a.WriteChapter(tt.args.chapterTitle, tt.args.content)
		})
	}
}

func TestAuthor_ReviewChapter(t *testing.T) {
	type fields struct {
		Name    string
		Contact string
	}
	type args struct {
		chapterTitle string
		content      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Author{
				Name:    tt.fields.Name,
				Contact: tt.fields.Contact,
			}
			a.ReviewChapter(tt.args.chapterTitle, tt.args.content)
		})
	}
}

func TestAuthor_FinalizeChapter(t *testing.T) {
	type fields struct {
		Name    string
		Contact string
	}
	type args struct {
		chapterTitle string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Author{
				Name:    tt.fields.Name,
				Contact: tt.fields.Contact,
			}
			a.FinalizeChapter(tt.args.chapterTitle)
		})
	}
}
