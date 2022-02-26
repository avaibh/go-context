package main

import (
	handler "avaibh/go-context/internal"
	"fmt"
	"net/http"
	"sync"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		server := createServer("Client", 9000, handler.Client)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	go func() {
		server := createServer("Middlewares", 9001, handler.Middleware)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	go func() {
		server := createServer("Server", 9002, handler.Server)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	wg.Wait()
}

func createServer(name string, port int, handler Handler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}
