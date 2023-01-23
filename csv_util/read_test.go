package csv_util

import "testing"

func TestRead(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"read test.csv",
			args{
				path: "test.csv",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Read(tt.args.path)
		})
	}
}
