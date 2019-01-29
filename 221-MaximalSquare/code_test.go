package leetcode

import (
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example",
			args{
				[][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			4,
		},
		{
			"boundary",
			args{
				[][]byte{
					{},
				},
			},
			0,
		},
		{
			"boundary",
			args{
				[][]byte{
					{'1', '1', '1', '1', '1'},
				},
			},
			1,
		},
		{
			"example",
			args{
				[][]byte{
					{'0', '0', '0', '1'},
					{'1', '1', '0', '1'},
					{'1', '1', '1', '1'},
					{'0', '1', '1', '1'},
					{'0', '1', '1', '1'},
				},
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
