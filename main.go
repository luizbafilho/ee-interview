package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()

	cache := NewCache()

	mux.HandleFunc("GET /users/{user}", fetchUserPublicGists(cache))

	// Handle Ctrl+C signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		fmt.Println("Server stopped")
		os.Exit(0)
	}()

	fmt.Println("Server listening on :8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}
