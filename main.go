package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	loginURL = "http://2.2.2.2/ac_portal/login.php"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}
	var (
		res string
		err error
	)
	switch args[0] {
	case "login":
		if len(args) < 3 {
			usage()
		}
		res, err = post("opr=pwdLogin&userName=" + args[1] + "&pwd=" + args[2] + "&remeberPwd=0")
	case "logout":
		res, err = post("opr=logout")
	default:
		usage()
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
func usage() {
	fmt.Println(`usage: ./login optype [stuid stupwd]
		optype can be login, logout
	`)
	os.Exit(1)
}

func post(data string) (string, error) {
	res, err := http.Post(loginURL, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
