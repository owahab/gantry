package main

import (
	"os"
	"os/exec"
	"fmt"
)

type Update struct {
}

func (u *Update) Run(args []string) error {
	fmt.Println("Updating registry...")
	return u.execute()
}

func (u *Update) runIfRequired() error {
	if u.isRequired() {
		fmt.Println("Local registry not found, downloading...")
		return u.execute()
	}
	return nil
}

func (u *Update) isRequired() bool {
	_, err := os.Stat(cacheDirectory())
	return err != nil
}

func (u *Update) execute() error {
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
	return cmd.Run()
}
