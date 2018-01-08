package codefights

import "bytes"

type Stack interface {
}

func reverseParenthesis(s string) string {
	return ""
}

func reverse(s string) string {
	var b bytes.Buffer

	for i := len(s) - 1; i >= 0; i-- {
		b.WriteRune(rune(s[i]))
	}

	return b.String()
}
