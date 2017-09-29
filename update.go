package main

import (
	"os"
	"os/exec"
)

func update() {
	directory := cacheDirectory()
	arguments := []string{}
	_, err := os.Stat(directory)
	if err == nil {
		arguments = append(arguments, "pull")
		arguments = append(arguments, "origin")
		arguments = append(arguments, "master")
	} else {
		arguments = append(arguments, "clone")
		arguments = append(arguments, Repository)
		arguments = append(arguments, directory)
	}
	cmd := exec.Command("git", arguments...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
