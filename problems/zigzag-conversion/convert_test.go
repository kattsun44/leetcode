package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "Case 2",
			args: args{s: "PAYPALISHIRING", numRows: 4},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "Case 3",
			args: args{s: "A", numRows: 1},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
