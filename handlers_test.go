package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchUserPublicGists(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/octocat", nil)
	require.NoError(t, err)
	req.SetPathValue("user", "octocat")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fetchUserPublicGists)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	// sily check of the octate id just to make sure we are getting a github response.
	expected := "583231"
	assert.Contains(t, rr.Body.String(), expected)
}
