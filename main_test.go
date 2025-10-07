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
	expectedStatus := http.StatusOK
	router := setupRouter()

	newRecord := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/healthcheck", nil)
	assert.NoError(t, err)

	expectedTime := time.Now()
	router.ServeHTTP(newRecord, req)

	var actualTime timeStruct
	err = json.Unmarshal(newRecord.Body.Bytes(), &actualTime)

	assert.Equal(t, expectedStatus, newRecord.Code)
	assert.Equal(t, "application/json; charset=utf-8", newRecord.Header().Get("Content-Type"))
	assert.NoError(t, err)
	assert.WithinDuration(t, expectedTime, actualTime.Time, time.Second)
}
