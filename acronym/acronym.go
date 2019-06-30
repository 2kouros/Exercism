package acronym

import "strings"

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	var st []string

	acro := ""

	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Join(strings.Fields(s), " ")
	st = strings.Split(strings.Title(s), " ")
	for _, v := range st {
		acro += string([]rune(v)[0])

	}
	return acro
}
