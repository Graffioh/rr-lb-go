package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var count = 0

func lb(w http.ResponseWriter, r *http.Request) {
	hosts := [3]string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"}

	count = (count + 1) % 3

	log.Printf(hosts[count])

	target, err := url.Parse(hosts[count])
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(w, r)
}

func createServer(port int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", lb)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func main() {
	http.HandleFunc("/", lb)

	log.Println("load balancer started")

	server := createServer(6969)
	server.ListenAndServe()
}
