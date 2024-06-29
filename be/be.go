package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from backend server")
}

func createServer(port int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	log.Println("backend servers started")

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
