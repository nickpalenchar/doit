package directives

import (
	"log"
	"os"
	"os/exec"
)

// In executes each shell command in the specified directory.
func In(directory string, commands []string) error {

	// Iterate over each command and execute it
	for _, cmdStr := range commands {
		cmd := exec.Command("sh", "-c", cmdStr)
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
