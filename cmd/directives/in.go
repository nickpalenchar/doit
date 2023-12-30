package directives

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// In executes each shell command in the specified directory.
func In(directory string, commands []string) error {

	if trimmed := strings.TrimSpace(directory); trimmed == "" {
		log.Fatalf("[ERROR] Directive IN must have a path (e.g. `IN .`; `IN /home`)")
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}

	// Iterate over each command and execute it
	for _, cmdStr := range commands {
		cmd := exec.Command(shell, "-c", cmdStr)
		cmd.Dir = directory
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("[ERROR] %s", err)
			return err
		}
	}

	// Change back to the original directory
	if err := os.Chdir(".."); err != nil {
		return err
	}

	return nil
}
