/*

Tasks:
- [ ] Implement each function by replacing its zero-value stub (marked TODO).
- [ ] Run `go test ./set01/` after every change.

Tools and techniques:
- Strings are immutable bytes; text characters are `rune`s. `for i, r := range s`
  walks a string rune-by-rune — this matters the instant you meet non-ASCII text.
- Skim https://pkg.go.dev/strings — knowing what a stdlib package offers is a
  core Go skill. Most "how do I..." questions are answered there.
- []rune(s) converts a string to runes; string(runes) converts back.

Expected time to complete: 45 minutes.

*/

package set01

import "strings"

// DONE FOR YOU. Note how the `strings` package is imported above and used as
// strings.<Name>.
func Shout(s string) string {
	return strings.ToUpper(s) + "!"
}

// CountVowels returns how many runes in s are a, e, i, o, or u (lowercase).
// Tip: strings.ContainsRune("aeiou", r) — or a switch.
func CountVowels(s string) int {
	return 0 // TODO: implement CountVowels
}

// Reverse returns s backwards. Reverse("hello") == "olleh".
// Convert to []rune first so "héllo" isn't mangled.
func Reverse(s string) string {
	return "" // TODO: implement Reverse
}

// IsPalindrome reports whether s reads the same both ways. Reuse Reverse.
func IsPalindrome(s string) bool {
	return false // TODO: implement IsPalindrome
}

// Initials returns the uppercased first rune of each word, joined together.
// Initials("grace brewster hopper") == "GBH"
// Tip: strings.Fields splits on whitespace; strings.ToUpper works on strings.
func Initials(s string) string {
	return "" // TODO: implement Initials
}
