package leetcode

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example",
			args{
				[]int{2, 7, 11, 15},
				9,
			},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}

	a := []byte(`this is a test`)
	fmt.Println(a)
	fmt.Println(string(a))
	fmt.Println(hex.EncodeToString(a))
	fmt.Println(base64.StdEncoding.EncodeToString(a))
}
