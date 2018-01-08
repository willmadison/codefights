package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreEquallyString(t *testing.T) {
	cases := []struct {
		given struct {
			left, right, friendsLeft, friendsRight int
		}
		expected bool
	}{
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{10, 15, 15, 10},
			true,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{15, 10, 15, 10},
			true,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{15, 10, 15, 9},
			false,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{10, 5, 5, 10},
			true,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{10, 15, 5, 20},
			false,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{10, 20, 10, 20},
			true,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{5, 20, 5, 20},
			true,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{20, 15, 5, 20},
			false,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{5, 10, 5, 10},
			true,
		},
		{
			struct {
				left, right, friendsLeft, friendsRight int
			}{1, 10, 10, 0},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := areEquallyStrong(tc.given.left, tc.given.right, tc.given.friendsLeft, tc.given.friendsRight)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
