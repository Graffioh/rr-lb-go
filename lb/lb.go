package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Server struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newServer(addr string) Server {
	serverurl, err := url.Parse(addr)
	if err != nil {
		log.Fatal("Error converting address into url")
	}

	return Server{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverurl),
	}
}

func lb(w http.ResponseWriter, r *http.Request, count *uint32, servers []Server) {
	idx := atomic.AddUint32(count, 1) % 3

	chosenServer := servers[idx]

	log.Printf("%v", chosenServer)

	chosenServer.proxy.ServeHTTP(w, r)
}

func createServer(port int, count *uint32, servers []Server) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb(w, r, count, servers)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func main() {
	var count uint32 = 0

	servers := []Server{
		newServer("http://localhost:8080"),
		newServer("http://localhost:8081"),
		newServer("http://localhost:8082"),
	}

	log.Println("load balancer started")

	server := createServer(6969, &count, servers)
	server.ListenAndServe()
}
