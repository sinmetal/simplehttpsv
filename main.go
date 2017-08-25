package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/compute/metadata"
)

type environment struct {
	Hostname string
	Zone     string
}

var environmentValue environment

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("path=%s, method=%s, header=%v", r.RequestURI, r.Method, r.Header)

	w.Header().Add("Cache-Control", "public,max-age=600")
	fmt.Fprintf(w, "version=1 %s, %s", environmentValue.Hostname, environmentValue.Zone)
}

func main() {
	hostname, err := metadata.Hostname()
	if err != nil {
		log.Fatalf("fatal get metadata.Hostname. %s", err.Error())
		return
	}

	zone, err := metadata.Zone()
	if err != nil {
		log.Fatalf("fatal get metadata.Zone. %s", err.Error())
		return
	}

	environmentValue.Hostname = hostname
	environmentValue.Zone = zone

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
