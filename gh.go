package main

import (
	"flag"
	"os"
)

const base = "https://api.github.com"

func init() {
	addHelp(
		"",
		`
gh is a simple tool that creates, lists and removes github repos.

Usage:

	gh command [arguments]

The commands are:

	login          prompts for and stores github username and password
	list   (ls)    lists all repos (grouped by owner)
	create (mk)    creates a repo
	delete (rm)    deletes a repo from github

Use "gh help [command]" for more information about a command.
		`,
	)
}

func main() {
	cmd := "gh"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "gh":
		cmdhelp("")
	case "h", "help":
		fs := flag.NewFlagSet("help", 0)
		fs.Parse(os.Args[2:])
		cmdhelp(fs.Arg(0))
	case "login":
		cmdlogin()
	case "ls", "list":
		cmdlist()
	case "mk", "create":
		fs := flag.NewFlagSet("create", 0)
		fPrivate := fs.Bool("p", false, "Make created repo private")
		fs.Parse(os.Args[2:])
		cmdcreate(fs.Arg(0), *fPrivate)
	case "rm", "delete":
		fs := flag.NewFlagSet("delete", 0)
		fs.Parse(os.Args[2:])
		cmddelete(fs.Arg(0))
	}
}
