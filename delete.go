package main

import (
	"log"
	"net/http"
)

func cmddelete(reponame string) {
	u, p, err := getCreds()
	if err != nil {
		log.Fatal(err)
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
