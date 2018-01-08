package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBorder(t *testing.T) {
	cases := []struct {
		given, expected []string
	}{
		{
			[]string{"abc", "ded"},
			[]string{"*****", "*abc*", "*ded*", "*****"},
		},

		{
			[]string{"a"},
			[]string{"***", "*a*", "***"},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := addBorder(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
