package main

import (
	"fmt"
	"net/http"
	"sync"

	handler "github.com/avaibh/go-context/internal"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		server := createServer("Client", 9000, handler.Client)
		server.ListenAndServe()
		wg.Done()
	}()

	go func() {
		server := createServer("Middleware", 9001, handler.Middleware)
		server.ListenAndServe()
		wg.Done()
	}()

	go func() {
		server := createServer("Server", 9002, handler.Server)
		server.ListenAndServe()
		wg.Done()
	}()

	wg.Wait()
}

func createServer(name string, port int, handler Handler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}
}
