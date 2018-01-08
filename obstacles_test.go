package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvoidObstacles(t *testing.T) {
	cases := []struct {
		given    []int
		expected int
	}{
		{
			[]int{5, 3, 6, 7, 9},
			4,
		},
		{
			[]int{2, 3},
			4,
		},
		{
			[]int{1, 4, 10, 6, 2},
			7,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := avoidObstacles(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
