package csv_util

import "testing"

func TestWrite(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"write test2.csv",
			args{path: "test2.csv"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Write(tt.args.path)
		})
	}
}
