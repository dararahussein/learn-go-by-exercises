/*

Files — reading and writing the disk.

Two workhorses:
  os.ReadFile(path)  -> ([]byte, error)   // whole file at once
  os.WriteFile(path, data, 0644) -> error // whole file at once
For line-by-line reading, a bufio.Scanner over an opened file is the idiom:

    f, err := os.Open(path)
    if err != nil { return nil, err }
    defer f.Close()                 // runs when the function returns
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()      // one line, no trailing newline
        ...
    }
    return ..., scanner.Err()

`defer` is new: it schedules f.Close() to run no matter how the function exits.

Tasks:
- [ ] Replace each panic("TODO...") with a real implementation.
- [ ] Run `go test ./set03/` after every change.

Expected time to complete: 45 minutes.

*/

package set03

import (
	"os"
	"strings"
)

// DONE FOR YOU: join the lines with newlines and write them out.
// 0644 is Unix file permissions (owner read/write, others read).
func WriteLines(path string, lines []string) error {
	data := strings.Join(lines, "\n")
	return os.WriteFile(path, []byte(data), 0644)
}

// ReadLines returns the file's lines with no trailing newlines.
// Use a bufio.Scanner as shown in the header (add the "bufio" import).
// An empty file should give an empty result and no error.
func ReadLines(path string) ([]string, error) {
	return nil, nil // TODO: implement ReadLines
}

// CountLines returns how many lines the file has.
// You already wrote ReadLines — reuse it.
func CountLines(path string) (int, error) {
	return 0, nil // TODO: implement CountLines
}

// WordFrequency reads the file and counts how often each whitespace-separated
// word appears. strings.Fields(s) splits on any run of whitespace.
// This is the classic "wc"-style tool, and the core of many real utilities.
func WordFrequency(path string) (map[string]int, error) {
	return nil, nil // TODO: implement WordFrequency
}
