package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func hello(w http.ResponseWriter, r *http.Request, port int) {
	fmt.Fprintf(w, "Hello from backend server on port %d", port)
}

func createServer(port int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hello(w, r, port)
	})

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

	for _, port := range []int{8080, 8081, 8082} {
		go func(p int) {
			server := createServer(p)
			server.ListenAndServe()
			wg.Done()
		}(port)
	}

	wg.Wait()
}
