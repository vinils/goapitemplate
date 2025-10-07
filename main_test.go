package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(test *testing.T) {
	time := time.Now()
	actual := newTime(time)
	expected := timeStruct{time}

	assert.Equal(test, actual, expected)
}

func TestNewTimeNow(test *testing.T) {
	actual := newTimeNow()
	expected := time.Now()

	assert.WithinDuration(test, expected, actual.Time, time.Second)
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	newRecord := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/healthcheck", nil)
	assert.NoError(t, err)

	expectedTime := time.Now()
	router.ServeHTTP(newRecord, req)

	assert.Equal(t, 200, newRecord.Code)
	// assert.Contains(t, newRecord.Body.String(), expected)

	// Assert the Content-Type header
	assert.Equal(t, "application/json; charset=utf-8", newRecord.Header().Get("Content-Type"))

	// Decode the response body
	var actualTime timeStruct
	err = json.Unmarshal(newRecord.Body.Bytes(), &actualTime)

	assert.NoError(t, err)
	assert.WithinDuration(t, expectedTime, actualTime.Time, time.Second)
}
