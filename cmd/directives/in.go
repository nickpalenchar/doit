package directives

import (
	"bytes"
	"fmt"
	"io"
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

func In(directory string, commands interface{}) error {
	if trimmed := strings.TrimSpace(directory); trimmed == "" {
		print.Error("Directive IN must have a path (e.g. `IN .`; `IN /home`)")
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}

	stdOut := &PrefixWriter{Prefix: ">> ", Writer: os.Stdout}
	stdErr := &PrefixWriter{Prefix: "!> ", Writer: os.Stderr}

	switch v := commands.(type) {
	case []interface{}:
		// Iterate over each command and execute it
		for _, cmdArg := range v {
			var cmdStr string
			var ignoreFailure bool

			switch cmd := cmdArg.(type) {
			case string:
				cmdStr = cmd
			case map[interface{}]interface{}:
				// Assume the map has a single key and value
				for k, v := range cmd {
					cmdStr = fmt.Sprintf("%v", v)
					if keyStr, ok := k.(string); ok && keyStr == "?" {
						ignoreFailure = true
					}
					break
				}
			default:
				print.Error("Invalid command type. Must be a string or a map with a single key.")
			}

			cmd := exec.Command(shell, "-c", cmdStr)
			cmd.Dir = directory
			cmd.Stdout = stdOut
			cmd.Stderr = stdErr
			print.Info("> " + cmdStr)
			if err := cmd.Run(); err != nil && !ignoreFailure {
				print.Error("check above error output (begins with !>)")
				print.Error(err.Error())
				print.Error("halting doit plan.")
				os.Exit(1)
			}
		}
	default:
		print.Error("Invalid command type. Must be a string or valid key/value pair.")
	}

	return nil
}
