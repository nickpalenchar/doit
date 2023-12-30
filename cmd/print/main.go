package print

import (
	"flag"
	"log"
	"os"
)

var (
	stdoutPrinter *log.Logger
	stderrPrinter *log.Logger
	verboseOutput bool
)

func init() {
	stdoutPrinter = log.New(os.Stdout, "", 0)
	stderrPrinter = log.New(os.Stderr, "", 0)
	flag.BoolVar(&verboseOutput, "v", false, "Enable verbose logging")
	flag.Parse()

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
