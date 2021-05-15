package find_minimum_in_rotated_sorted_array

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{3, 4, 5, 1, 2}}, want: 1},
		{name: "test2", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{3, 4, 5, 1, 2}}, want: 1},
		{name: "test2", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}}, want: 0},
		{name: "test3", args: args{nums: []int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}}, want: 0},
		{name: "test4", args: args{nums: []int{11, 13, 15, 17}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin2(tt.args.nums); got != tt.want {
				t.Errorf("findMin2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{3, 4, 5, 1, 2}}, want: 1},
		{name: "test2", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}}, want: 0},
		{name: "test3", args: args{nums: []int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}}, want: 0},
		{name: "test4", args: args{nums: []int{11, 13, 15, 17}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin3(tt.args.nums); got != tt.want {
				t.Errorf("findMin3() = %v, want %v", got, tt.want)
			}
		})
	}
}
