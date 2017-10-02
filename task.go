package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"os"
	"os/exec"
	"path/filepath"
)

type Task struct {
	options     []string
	Name        string   `yaml:"name"`
	Image       string   `yaml:"image"`
	Workdir     string   `yaml:"workdir"`
	Volumes     []string `yaml:"volumes"`
	Environment []string `yaml:"environment"`
}

func (t *Task) Run(args []string) error {
	if len(args) == 0 {
		fmt.Println("Command too short")
		return nil
	}

	t.readConfig(args[0])

	if len(t.Image) == 0 {
		fmt.Println(fmt.Sprintf("Unknown command: %q", args[0]))
		return nil
	}

	t.options = append(t.options, "run")
	t.options = append(t.options, "--rm")
	t.options = append(t.options, "--interactive")

	t.addOptional("workdir", t.Workdir)
	if len(t.Volumes) > 0 {
		for _, vol := range t.Volumes {
			t.addOptional("volume", vol)
		}
	}
	if len(t.Environment) > 0 {
		for _, env := range t.Environment {
			t.addOptional("env", env)
		}
	}
	t.options = append(t.options, t.Image)

	// Merge with arguments
	t.options = append(t.options, args...)
	t.execute()
	return nil
}

func (t *Task) readConfig(command string) {
	file := filepath.Join(cacheDirectory(), fmt.Sprintf("%v.yml", command))
	yamlFile, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(yamlFile, t)
}

func (t *Task) addOptional(name string, value string) {
	if value != "" {
		value = os.ExpandEnv(value)
		t.options = append(t.options, fmt.Sprintf("--%s=%s", name, strings.TrimSpace(value)))
	}
}

func (t *Task) execute() {
	pwd, _ := os.Getwd()
	cmd := exec.Command("docker", t.options...)
	cmd.Dir = pwd
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	fmt.Println(fmt.Sprintf("Running command: %q", strings.Join(cmd.Args, " ")))
	cmd.Run()
}
