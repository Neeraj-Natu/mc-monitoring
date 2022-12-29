package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Host struct {
	HostName string `json:"HostName"`
	Cluster  string `json:"Cluster"`
	Service  string `json:"Service"`
}

var cluster string
var service string

func serveHTTP(w http.ResponseWriter, r *http.Request) {

	Hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	var Host = Host{HostName: Hostname, Cluster: cluster, Service: service}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Host)
}

func main() {

	cluster = os.Args[1]
	service = os.Args[2]
	http.HandleFunc("/", serveHTTP)
	err := http.ListenAndServe(":9090", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
