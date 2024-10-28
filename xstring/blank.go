package xstring

import "strings"

// IsBlank Determine if the string is blank, like " ", "\t", "\n"
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
