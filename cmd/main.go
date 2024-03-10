package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/nickpalenchar/doit/cmd/directives"
	"github.com/nickpalenchar/doit/cmd/print"
	"gopkg.in/yaml.v2"
)

var Version string
var (
	showVersion   bool
	verboseOutput bool
)

func getVersion() string {
	if Version == "" {
		Version = "0.0.0-dev"
	}
	return Version
}

type Config struct {
	Main yaml.MapSlice `yaml:"__MAIN__"`
}

func init() {
	flag.BoolVar(&showVersion, "version", false, "Print version and exit")
	flag.BoolVar(&verboseOutput, "v", false, "Use verbose logging")
	flag.Parse()

	if showVersion {
		fmt.Printf("Version: %s\n", getVersion())
		os.Exit(0)
	}
	print.SetVerboseOutput(verboseOutput)

	print.Debug("verbose logging activated")

}

func stripFirstWord(input string) string {
	re := regexp.MustCompile(`^[^\s]+\s*`)
	result := re.ReplaceAllString(input, "")
	return result
}

func executeDirective(directive string, commands interface{}) {
	if strings.HasPrefix(directive, "IN") {
		directory := stripFirstWord(directive)
		print.Debug(fmt.Sprintf("Starting commands in %s", directory))
		directives.In(directory, commands)
	}
}

func runCommandsFromYAML(yamlFile string) {
	config := Config{}

	file, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	for _, item := range config.Main {
		if directive, ok := item.Key.(string); ok {
			executeDirective(directive, item.Value)
		} else {
			log.Fatalf("Directive '%v' is invalid.", directive)
		}
	}
}

func findYAMLFile() (string, error) {
	yamlFiles := []string{"doit.yml", "doit.yaml"}

	for _, file := range yamlFiles {
		_, err := os.Stat(file)
		if err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("No 'doit.yml' or 'doit.yaml' file found")
}

func main() {
	yamlFile, err := findYAMLFile()
	print.Debug(fmt.Sprintf("Loaded yaml file: %s", yamlFile))
	if err != nil {
		log.Fatal(err)
	}
	runCommandsFromYAML(yamlFile)
}
