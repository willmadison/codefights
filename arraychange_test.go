package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayChange(t *testing.T) {
	cases := []struct {
		given    []int
		expected int
	}{
		{
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{-1000, 0, -2, 0},
			5,
		},
		{
			[]int{2, 1, 10, 1},
			12,
		},
		{
			[]int{2, 3, 3, 5, 5, 5, 4, 12, 12, 10, 15},
			13,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := arrayChange(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
