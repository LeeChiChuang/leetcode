package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{19,1,3,49,4}}, []int{1,3,4,19,49}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, QuickSort(tt.args.nums), "they should be equal")
	}
}
