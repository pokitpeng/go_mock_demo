package utils

import "testing"

func TestHasString(t *testing.T) {
	type args struct {
		array []string
		elem  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{array: []string{"he", "ll", "ow"}, elem: "he"}, true},
		{"", args{array: []string{"he", "ll", "ow"}, elem: "ho"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasString(tt.args.array, tt.args.elem); got != tt.want {
				t.Errorf("HasString() = %v, want %v", got, tt.want)
			}
		})
	}
}
