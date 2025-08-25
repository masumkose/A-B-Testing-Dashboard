package main

import (
	"bytes"
	"encoding/json"
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

	database.Migrator().DropTable(&Experiment{}, &Variation{})

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


func TestAssignmentAndConversionEndpoints(t *testing.T) {
	// 1. Set up a clean database and router.
	setupTestDB()
	router := setupRouter()

	// 2. First, create an experiment to work with.
	expPayload := []byte(`{"name": "Core Logic Test", "variations": ["Control", "Candidate"]}`)
	req, _ := http.NewRequest("POST", "/api/experiments", bytes.NewBuffer(expPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	
	var createdExperiment Experiment
	json.Unmarshal(w.Body.Bytes(), &createdExperiment)
	assert.Equal(t, "Core Logic Test", createdExperiment.Name)
	assert.Equal(t, 2, len(createdExperiment.Variations))

	// 3. Test the Assignment Endpoint
	t.Run("Assign User", func(t *testing.T) {
		assignReq, _ := http.NewRequest("GET", "/api/experiments/1/assign", nil)
		assignW := httptest.NewRecorder()
		router.ServeHTTP(assignW, assignReq)

		assert.Equal(t, http.StatusOK, assignW.Code)

		var jsonResponse map[string]interface{}
		json.Unmarshal(assignW.Body.Bytes(), &jsonResponse)
		
		// Check that the response contains one of the valid variation names.
		variationName := jsonResponse["variationName"].(string)
		assert.Contains(t, []string{"Control", "Candidate"}, variationName)

		// Check that the participant count was incremented in the database.
		var updatedExp Experiment
		DB.Preload("Variations").First(&updatedExp, 1)
		totalParticipants := updatedExp.Variations[0].Participants + updatedExp.Variations[1].Participants
		assert.Equal(t, uint(1), totalParticipants)
	})

	// 4. Test the Conversion Endpoint
	t.Run("Record Conversion", func(t *testing.T) {
		// Let's assume the user was assigned variation with ID 2 ("Candidate").
		convertReq, _ := http.NewRequest("POST", "/api/variations/2/convert", nil)
		convertW := httptest.NewRecorder()
		router.ServeHTTP(convertW, convertReq)

		assert.Equal(t, http.StatusOK, convertW.Code)
		assert.Contains(t, convertW.Body.String(), "Conversion recorded")

		// Check that the conversion count was incremented in the database.
		var convertedVariation Variation
		DB.First(&convertedVariation, 2)
		assert.Equal(t, uint(1), convertedVariation.Conversions)
	})
}