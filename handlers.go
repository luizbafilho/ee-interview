package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchUserPublicGists(cache cacher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.PathValue("user")

		cacheKey := fmt.Sprintf("user:%s:gists", user)
		cachedResponse, ok := cache.Get(cacheKey)
		if ok {
			w.Write([]byte(cachedResponse))
			return
		}

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
	}
}
