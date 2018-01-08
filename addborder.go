package codefights

import (
	"fmt"
	"strings"
)

func addBorder(picture []string) []string {
	var bordered []string
	headerAndFooter := strings.Repeat("*", len(picture[0])+2)

	bordered = append(bordered, headerAndFooter)

	for _, v := range picture {
		bordered = append(bordered, fmt.Sprintf("*%s*", v))
	}

	bordered = append(bordered, headerAndFooter)

	return bordered
}
