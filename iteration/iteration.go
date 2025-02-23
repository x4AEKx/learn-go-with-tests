package iteration

import "strings"

func Repeat(value string) string {
	var repeated strings.Builder

	for i := 0; i < 5; i++ {
		repeated.WriteString(value)
	}

	return repeated.String()
}
