package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

func lb(w http.ResponseWriter, r *http.Request, count *uint32) {
	hosts := [3]string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"}

	idx := atomic.AddUint32(count, 1) % 3

	log.Printf(hosts[idx])

	target, err := url.Parse(hosts[idx])
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(w, r)
}

func createServer(port int, count *uint32) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb(w, r, count)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func main() {
	var count uint32 = 0

	log.Println("load balancer started")

	server := createServer(6969, &count)
	server.ListenAndServe()
}
