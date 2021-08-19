package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var server = "localhost:8080"

func main() {
	fmt.Println("Start app")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func hello(rw http.ResponseWriter, rq *http.Request) {
	if rq.RequestURI == "/terminate" {
		os.Exit(0)
	} else if rq.RequestURI != "/favicon.ico" && rq.RequestURI != "/" {
		_, err := fmt.Fprintf(rw, "hello %s", strings.Replace(rq.RequestURI, "/", "", 1))
		if err != nil {
			return
		}
	} else {
		_, err := fmt.Fprintf(rw, "Add your name to URI. For example: %s/myName", server)
		if err != nil {
			return
		}
	}
}
