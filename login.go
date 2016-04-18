package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func cmdlogin() {
	_, _, err := getCreds()
	if err != nil {
		log.Fatal(err)
	}
}

func getCreds() (username, password string, err error) {
	u, p, err := cache()
	if err != nil {
		return "", "", err
	}
	if u != "" && p != "" {
		return u, p, nil
	}
	u, p, err = prompt()
	if err != nil {
		return "", "", err
	}
	err = storeCreds(u, p)
	if err != nil {
		log.Fatal(err)
	}
	return u, p, nil
}

func cache() (username, password string, err error) {
	u, err := user.Current()
	if err != nil {
		return "", "", err
	}
	b, err := ioutil.ReadFile(u.HomeDir + "/.gh")
	if err != nil {
		return "", "", nil
	}
	l := strings.Split(string(b), "\n")
	if len(l) < 2 {
		return "", "", errors.New("malformed .gh file")
	}
	return l[0], l[1], nil
}

func prompt() (username, password string, err error) {
	fmt.Print("Github username: ")
	if _, err = fmt.Scanln(&username); err != nil {
		return "", "", err
	}

	fmt.Print("password: ")
	p, err := terminal.ReadPassword(0)
	if err != nil {
		return "", "", err
	}
	fmt.Println("")
	return username, string(p), nil
}

func storeCreds(username, password string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	contents := []byte(username + "\n" + password)
	return ioutil.WriteFile(u.HomeDir+"/.gh", contents, 0644)
}
