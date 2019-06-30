// Package bob answers based on the input statement
package bob

import (
	"strings"
)

// Hey provides a response based on input string characteristics
func Hey(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "Fine. Be that way!"
	}

	question := strings.HasSuffix(s, "?")
	upper := strings.ToUpper(s)
	lower := strings.ToLower(s)

	if question {
		if s == upper && s != lower {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if upper == s && lower != s {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
