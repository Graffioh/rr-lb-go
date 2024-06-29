package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "\nHello from Backend Server")
	fmt.Fprint(w, sb.String())
}

func createServer(port int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	http.HandleFunc("/", handler)

	log.Println("Backend Servers ON")

	go func() {
		server := createServer(8080)
		server.ListenAndServe()

		wg.Done()
	}()

	go func() {
		server := createServer(8081)
		server.ListenAndServe()

		wg.Done()
	}()

	go func() {
		server := createServer(8082)
		server.ListenAndServe()

		wg.Done()
	}()

	wg.Wait()
}
