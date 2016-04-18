package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	addHelp(
		"create",
		`
gh create (shortname: gh mk) creates a github repo.

Usage: 
	
	gh create reponame

git setup refresher:

	git init
	git remote add origin https://github.com/username/reponame
	git push --set-upstream origin master
		`,
	)
	addHelp(
		"mk",
		`
gh mk (longname: gh create) creates a github repo.

Usage: 
	
	gh mk reponame

git setup refresher:

	git init
	git remote add origin https://github.com/username/reponame
	git push --set-upstream origin master
		`,
	)
}

func cmdcreate(reponame string) {
	b, err := json.Marshal(&struct {
		Name string `json:"name"`
	}{
		Name: reponame,
	})
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", base+"/user/repos", bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	u, p, err := getCreds()
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(u, p)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 201 {
		log.Fatal("POST /user/repos: status: ", resp.StatusCode)
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
}
