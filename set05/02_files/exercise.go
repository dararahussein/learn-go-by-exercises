/* Files introduce resource ownership: open, check, defer Close, scan, check. */
package files

import (
	"os"
	"strings"
)

func WriteLines(path string, lines []string) error {
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644)
}

func ReadLines(path string) ([]string, error) {
	return nil, nil // TODO: os.Open, defer Close, bufio.Scanner, scanner.Err
}

func CountLines(path string) (int, error) {
	return 0, nil // TODO: reuse ReadLines
}

func WordFrequency(path string) (map[string]int, error) {
	return nil, nil // TODO: reuse ReadLines and strings.Fields
}
