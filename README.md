# Gantry

Gantry is a command runner tool that uses Docker containers to run your commands.

No more compiling, installing dependencies or altering configuration.

Gantry downloads and maintains a database of commands from here: https://github.com/owahab/gantry-registry

You can alter the registry on your machine to your liking.

## Install

### Windows

Download the binaries: https://github.com/owahab/gantry/releases/latest

### Linux

Download the binaries: https://github.com/owahab/gantry/releases/latest

### Mac

If you have Homebrew, just run:

    $ brew install owahab/tap/gantry

## Usage

Before running your first command, make sure you have updated your local command cache.

Run:

    $ gantry update

Now start running commands in containers:

    $ gantry run <command> <parameters>
