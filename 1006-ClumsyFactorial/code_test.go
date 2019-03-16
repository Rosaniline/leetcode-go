package leetcode

import "testing"

func Test_clumsy(t *testing.T) {
	tests := []struct {
		input1 int
		want   int
	}{
		{4, 7},
		{10, 12},
		{1, 1},
		{2, 2},
		{3, 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := clumsy(tt.input1); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
