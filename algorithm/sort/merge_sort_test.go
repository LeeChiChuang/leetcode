package sort

import (
"github.com/stretchr/testify/assert"
"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test2", args{[]int{19,1,3,49,4}}, []int{1,3,4,19,49}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MergeSort(tt.args.nums), "they should be equal")
	}
}

