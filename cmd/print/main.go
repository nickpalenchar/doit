package print

import (
	"log"
	"os"
)

var (
	stdoutPrinter *log.Logger
	stderrPrinter *log.Logger
)
var verboseOutput = false

func init() {
	stdoutPrinter = log.New(os.Stdout, "", 0)
	stderrPrinter = log.New(os.Stderr, "", 0)

}

func SetVerboseOutput(useVerboseOutput bool) {
	verboseOutput = useVerboseOutput
}

func Debug(msg string) {
	if verboseOutput {
		stdoutPrinter.Printf("[DEBUG] %s", msg)
	}
}

func Info(msg string) {
	stdoutPrinter.Printf(msg)
}

func Error(msg string) {
	stderrPrinter.Printf("[ERROR] %s", msg)
}
