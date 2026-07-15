package textutil

import "strings"

// NormalizeISBN removes separators and normalizes case for comparisons.
func NormalizeISBN(s string) string {
	r := strings.NewReplacer("-", "", " ", "")
	return strings.ToUpper(r.Replace(strings.TrimSpace(s)))
}
