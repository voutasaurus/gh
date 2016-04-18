package main

import "flag"

const base = "https://api.github.com"

func main() {
	flag.Parse()
	switch flag.Arg(0) {
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
