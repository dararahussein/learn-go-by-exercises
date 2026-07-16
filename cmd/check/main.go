package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type event struct {
	Action string
	Output string
	Test   string
}

var (
	testMessage = regexp.MustCompile(`^\s+\S+_test\.go:\d+:\s+(.*)$`)
	comparison  = regexp.MustCompile(`^(.*): got (.*), want (.*)$`)
)

func main() {
	cmd := exec.Command("go", "test", "-json")
	output, err := cmd.CombinedOutput()

	message, compilerOutput := readOutput(output)
	switch {
	case message != "":
		printMessage(message)
	case err == nil:
		fmt.Println("All exercises pass.")
	case compilerOutput != "":
		fmt.Print(compilerOutput)
	default:
		fmt.Println("The tests could not run. Try `go test` for full details.")
	}

	if err != nil {
		os.Exit(1)
	}
}

func readOutput(output []byte) (string, string) {
	decoder := json.NewDecoder(bytes.NewReader(output))
	var compilerLines []string

	for decoder.More() {
		var e event
		if err := decoder.Decode(&e); err != nil {
			break
		}
		line := strings.TrimSuffix(e.Output, "\n")
		if match := testMessage.FindStringSubmatch(line); match != nil {
			return match[1], ""
		}
		if e.Test == "" && isCompilerMessage(line) {
			compilerLines = append(compilerLines, line)
		}
	}

	return "", strings.Join(compilerLines, "\n") + newlineIfNeeded(compilerLines)
}

func isCompilerMessage(line string) bool {
	trimmed := strings.TrimSpace(line)
	return trimmed != "" &&
		trimmed != "FAIL" &&
		trimmed != "PASS" &&
		!strings.HasPrefix(trimmed, "exit status") &&
		!strings.HasPrefix(trimmed, "FAIL\t") &&
		!strings.HasPrefix(trimmed, "ok\t") &&
		!strings.HasPrefix(trimmed, "# ")
}

func newlineIfNeeded(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return "\n"
}

func printMessage(message string) {
	match := comparison.FindStringSubmatch(message)
	if match == nil {
		fmt.Println(message)
		return
	}

	fmt.Println(match[1])
	fmt.Printf("  your output: %s\n", match[2])
	fmt.Printf("  expected:    %s\n", match[3])
}
