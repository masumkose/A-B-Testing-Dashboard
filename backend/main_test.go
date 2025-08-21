package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	// set up our router for the test
	router := setupRouter()

	// httptest.NewRecorder() creates a "ResponseRecorder", which acts like a fake web browser
	// that captures the response from our server

	w := httptest.NewRecorder()

	// http.NewRequest() creates a fake request to our "/ping" endpoint.
	req, _ := http.NewRequest("GET", "/ping", nil)

	// router.ServeHTTP() sends our fake request to our router and recors the response in "w".
	router.ServeHTTP(w, req)

	// assert.Equal() is from the testify library. It checks if two values are equal.
	// This is much cleaner than writing if/else blocks for tests.
	// Here, we check if the HTTP status code in the response is 200 (OK).
	assert.Equal(t, http.StatusOK, w.Code)

	// Here, we check if the body of the response is the exact string we expect.
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())

}
