package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/{user}", func(w http.ResponseWriter, r *http.Request) {
		user := r.PathValue("user")

		// We can use the query parameters from the original request
		// so we get pagination and other options out of the box.
		queryParams := r.URL.Query()

		resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists?%s", user, queryParams.Encode()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusNotFound {
				http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
				return
			}

			http.Error(w, `{"error": "Failed to retrieve gists"}`, http.StatusInternalServerError)
			return
		}

		// Given that isn't a requirement to make any changes to the response, we can just copy the response body
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
	})

	// Handle Ctrl+C signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		fmt.Println("Server stopped")
		os.Exit(0)
	}()

	fmt.Println("Server listening on :8080")
	http.ListenAndServe("127.0.0.1:8080", mux)
}
