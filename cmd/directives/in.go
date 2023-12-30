package directives

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/nickpalenchar/doit/cmd/print"
)

type PrefixWriter struct {
	Prefix string
	Writer io.Writer
}

func (pw *PrefixWriter) Write(p []byte) (n int, err error) {
	lines := bytes.Split(p, []byte("\n"))
	for _, line := range lines {
		if len(line) > 0 {
			// Add the prefix and write the modified line
			n, err = pw.Writer.Write(append([]byte(pw.Prefix), line...))
			if err != nil {
				return n, err
			}
		}

		// Add newline after each line
		n, err = pw.Writer.Write([]byte("\n"))
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}

// In executes each shell command in the specified directory.
func In(directory string, commands []string) error {

	if trimmed := strings.TrimSpace(directory); trimmed == "" {
		log.Fatalf("[ERROR] Directive IN must have a path (e.g. `IN .`; `IN /home`)")
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}

	stdOut := &PrefixWriter{Prefix: ">> ", Writer: os.Stdout}

	// Iterate over each command and execute it
	for _, cmdStr := range commands {
		cmd := exec.Command(shell, "-c", cmdStr)
		cmd.Dir = directory
		cmd.Stdout = stdOut
		cmd.Stderr = os.Stderr
		print.Info("> " + cmdStr)
		if err := cmd.Run(); err != nil {
			log.Fatalf("[ERROR] %s", err)
			return err
		}
	}
	return nil
}
