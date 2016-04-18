package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	addHelp(
		"delete",
		`
gh delete (shortname: gh rm) removes a github repo.

Usage: 
	
	gh delete reponame

Warning:

	This action is irreversible. Any data in the remote repo that is not backed
	up will be lost forever. 
		`,
	)
	addHelp(
		"rm",
		`
gh rm (longname: gh delete) removes a github repo.

Usage: 
	
	gh rm reponame

Warning:

	This action is irreversible. Any data in the remote repo that is not backed
	up will be lost forever.
		`,
	)
}

func cmddelete(reponame string) {
	u, p, err := getCreds()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"You are about to delete the remote repository at %q\n",
		"github.com/"+u+"/"+reponame,
	)
	fmt.Printf("Are you sure? (Y/N):")
	if abort() {
		return
	}

	endpoint := base + "/repos/" + u + "/" + reponame
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(u, p)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 204 {
		log.Fatal("DELETE ", endpoint, ": status: ", resp.StatusCode)
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
}

func abort() bool {
	var confirm string
	for {
		fmt.Scanln(&confirm)
		switch confirm {
		case "y", "Y", "yes", "YES", "Yes":
			return false
		case "n", "N", "no", "NO", "No":
			return true
		default:
			fmt.Println("please type Y to delete, N to abort")
		}
	}
}
