package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"test1", args{"aab"}, [][]string{{"a","a","b"}, {"aa", "b"}}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Partition(tt.args.s), "they should be equal")
	}
}