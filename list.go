package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

func init() {
	addHelp(
		"list",
		`
gh list (shortname: gh ls) lists all github repos the user has write access to.

Usage: 
	
	gh list

listed repos are grouped by owner and sorted alphabetically.
		`,
	)
	addHelp(
		"ls",
		`
gh ls (longname: gh list) lists all github repos the user has write access to.

Usage: 
	
	gh ls

listed repos are grouped by owner and sorted alphabetically.
		`,
	)
}

func cmdlist() {
	u, p, err := getCreds()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", base+"/user/repos", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(u, p)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("GET /user/repos: status: ", resp.StatusCode)
	}
	var x []struct {
		Owner struct {
			Login string
		}
		Name string
	}
	err = json.NewDecoder(resp.Body).Decode(&x)
	if err != nil {
		log.Fatal(err)
	}
	byowner := make(map[string][]string)
	for _, repo := range x {
		byowner[repo.Owner.Login] = append(byowner[repo.Owner.Login], repo.Name)
	}
	owners := []string{}
	for o := range byowner {
		owners = append(owners, o)
	}
	sort.Strings(owners)
	for _, o := range owners {
		fmt.Println(o + ":")
		for _, r := range byowner[o] {
			fmt.Println("  " + r)
		}
	}
}
