package main

import (
	"os"
	"os/exec"
	"fmt"
)

type Update struct {
}

func (u *Update) Run() {
	fmt.Println("Updating registry...")
}

func (u *Update) RunIfRequired()  {
	if u.IsRequired() {
		fmt.Println("Local registry not found, downloading...")
		u.execute()
	}
}

func (u *Update) IsRequired() bool {
	_, err := os.Stat(cacheDirectory())
	if err != nil {
		return true
	}

	return false
}

func (u *Update) execute() {
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
