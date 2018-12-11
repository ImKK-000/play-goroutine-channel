package main

import (
	"fmt"
	"log"
	"net/http"
)

var channel = make(chan string, 10)
var ports = []int{8000, 8001, 8002}

type APIHandler struct{}

func (apiHandler APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	channel <- r.Host
	w.Write([]byte("done\n"))
}

func NewServer(port int) {
	apiHandler := APIHandler{}
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), apiHandler)
	log.Fatalln(err)
}

func main() {
	go NewServer(ports[0])
	go NewServer(ports[1])

	for {
		fmt.Println(<-channel)
		fmt.Println("??")
	}
}
