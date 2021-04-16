package next_greater_element_i

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement2(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			[]int{-1, 3, -1},
		},
		{
			"case 2",
			args{
				nums1: []int{2,4},
				nums2: []int{1,2,3,4},
			},
			[]int{3,-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement2(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			[]int{-1, 3, -1},
		},
		{
			"case 2",
			args{
				nums1: []int{2,4},
				nums2: []int{1,2,3,4},
			},
			[]int{3,-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
