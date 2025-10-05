package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func diffTimeAndMaxDiffSeconds(time time.Time, actual timeStruct) (float64, float64) {
	diff := time.Sub(actual.Time).Seconds()
	maxDiffSeconds := float64(10)
	return diff, maxDiffSeconds
}

func TestNewTime(test *testing.T) {
	time := time.Now()
	expected := timeStruct{time}
	actual := newTime(time)

	assert.Equal(test, actual, expected)
}

func TestNewTimeNow(test *testing.T) {
	expected := time.Now()
	actual := newTimeNow()
	difTime, maxDiffSeconds := diffTimeAndMaxDiffSeconds(expected, actual)

	assert.LessOrEqual(test, difTime, maxDiffSeconds)
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

	difTime, maxDiffSeconds := diffTimeAndMaxDiffSeconds(expectedTime, actualTime)

	assert.LessOrEqual(t, difTime, maxDiffSeconds)
}

func TestGetenvOrDefault_WhenNoValue(test *testing.T) {
	expected := "defaultvalue"
	actual := GetenvOrDefault("", expected)

	assert.Equal(test, actual, expected)
}

func TestGetenvOrDefault_EnvValue(test *testing.T) {
	key := "keyEnviromentName"
	os.Setenv(key, "expectedvalue")
	expected := os.Getenv(key)
	actual := GetenvOrDefault(key, "notexpected")

	assert.Equal(test, actual, expected)
}
