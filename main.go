package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
	"runtime"
	"path/filepath"
)

const Name = "gantry"
const Version = "0.4.1"
const Usage = "Run Commands inside Docker Containers"
const Repository = "https://github.com/docker-gantry/registry"
const LocalCacheDirectory = ".gantry"

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = Usage

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run a command",
			Action: func(c *cli.Context) error {
				var t Task
				if err := t.Run(c.Args()); err != nil {
					return cli.NewExitError(err.Error(), 2)
				}
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   fmt.Sprintf("update %v", Name),
			Action: func(c *cli.Context) error {
				var u Update
				if err := u.Run(c.Args()); err != nil {
					return cli.NewExitError(err.Error(), 2)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func homeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func cacheDirectory() string {
	return filepath.Join(homeDir(), LocalCacheDirectory)
}