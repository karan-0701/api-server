package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for "Hello World" GET endpoint
func TestHelloWorldShouldRun(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	// returns byte string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

// Test for invalid "Hello World" POST request
func TestHelloWorldWouldNotRun(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("Some Body")

	fmt.Println(testServer.URL)
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode) // Assuming POST is not allowed for this endpoint
}

// Test for health GET endpoint
func TestHealthShouldRun(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	// returns byte string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Health", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

// Test for invalid health POST request
func TestHealthWouldNotRun(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("Some Body")

	fmt.Println(testServer.URL)
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

// Test for endpoint GET request
func TestEndpointShouldRun(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleEndPoint))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	// returns byte string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "new EndPoint", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

// Test for invalid endpoint POST request
func TestEndpointWouldNotRun(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleEndPoint))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("Some Body")

	fmt.Println(testServer.URL)
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
