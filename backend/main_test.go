package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	// "file:memeory:" is the special connection string for an in-memory SQLite DB.
	// "?cache=shared" allows the connection to be shared across the test function.

	database, err := gorm.Open(sqlite.Open("file::memory?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database!")
	}

	err = database.AutoMigrate(&Experiment{}, &Variation{})
	if err != nil {
		panic("Failed to migrate test database!")
	}

	DB = database
}

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

func TestExperimentEndpoints(t *testing.T) {
	// 1. Set up a clean database and a new router for this specific test.
	setupTestDB()
	router := setupRouter()

	t.Run("Create Experiment Success", func(t *testing.T) {
		// create a json payload for our request body
		payload	:= []byte(`{"name": "Test Experiment", "variations": ["A", "B"]}`)
		req, _	:= http.NewRequest("POST", "/api/experiments", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		assert.Contains(t, w.Body.String(), `"Name":"Test Experiment"`)
		assert.Contains(t, w.Body.String(), `"Name":"A"`)

	})

	t.Run("Get Experiments", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/experiments", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, w.Body.String(), `"Name":"Test Experiment"`)


		var experiments []Experiment
		DB.Find(&experiments)
		assert.Equal(t, 1, len(experiments))
		assert.Equal(t, "Test Experiment", experiments[0].Name)
	})



}
