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
	options   []string
	Name      string `yaml:"name"`
	Image     string `yaml:"image"`
	Volume    string `yaml:"volume"`
	Workdir   string `yaml:"workdir"`
}

func (t *Task) Run(args []string) error {
	if len(args) == 0 {
		fmt.Println("Command too short")
		return nil
	}

	t.readConfig(args[0])

	t.options = append(t.options, "run")
	t.addOptional("volume", t.Volume)
	t.addOptional("workdir", t.Workdir)

	//t.options = append(t.options, "--rm")
	t.options = append(t.options, "--interactive")

	// Merge with arguments
	t.options = append(t.options, args...)
	t.execute()
	return nil
}

func (t *Task) readConfig(command string) {
	file := fmt.Sprintf(filepath.Join(cacheDirectory(), "%v.yml", command))
	yamlFile, _ := ioutil.ReadFile(file)
	yaml.Unmarshal(yamlFile, t)
}

func (t *Task) addOptional(name string, value string) {
	if value != "" {
		value = os.ExpandEnv(value)
		t.options = append(t.options, fmt.Sprintf("--%v=%v", name, strings.TrimSpace(value)))
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
	cmd.Run()
}