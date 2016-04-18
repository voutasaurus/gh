package main

import "flag"

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
	flag.Parse()
	switch flag.Arg(0) {
	case "h", "help":
		cmdhelp(flag.Arg(1))
	case "login":
		cmdlogin()
	case "ls", "list":
		cmdlist()
	case "mk", "create":
		cmdcreate(flag.Arg(1))
	case "rm", "delete":
		cmddelete(flag.Arg(1))
	}
}
