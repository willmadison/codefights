package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayMaximalAdjacentDifference(t *testing.T) {
	cases := []struct {
		given    []int
		expected int
	}{
		{
			[]int{2, 4, 1, 0},
			3,
		},
		{
			[]int{1, 1, 1, 1},
			0,
		},
		{
			[]int{-1, 4, 10, 3, -2},
			7,
		},
		{
			[]int{-1, 4, 10, 3, -2},
			7,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := arrayMaximalAdjacentDifference(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
