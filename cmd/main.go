package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/nickpalenchar/doit/cmd/directives"
	"github.com/nickpalenchar/doit/cmd/print"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Main map[string]interface{} `yaml:"__MAIN__"`
}

var verbose bool

func init() {
	// flag.BoolVar(&verbose, "v", false, "Enable verbose logging")
	// flag.Parse()
	print.Debug("verbose logging activated")
}

func stripFirstWord(input string) string {
	// Define a regular expression to match the first word followed by spaces
	re := regexp.MustCompile(`^[^\s]+\s*`)

	// Replace the matched part with an empty string
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

	for directive, args := range config.Main {
		executeDirective(directive, args)
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
