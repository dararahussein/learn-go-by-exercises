package set04

import "strings"

// NormalizeSpaces should trim the input and replace every run of whitespace
// (spaces, tabs, and newlines) with one ordinary space.
//
// BUG: the current implementation handles one narrow example and nothing else.
func NormalizeSpaces(s string) string {
	return strings.ReplaceAll(s, "  ", " ")
}
