//go:build integration
// +build integration

package rest_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

// TestPingHandlerIntegration tests the ping handler in a real HTTP environment
func TestPingHandlerIntegration(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	// Configuration
	baseURL := getTestBaseURL()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	t.Run("Ping Endpoint Health Check", func(t *testing.T) {
		// Make request to ping endpoint
		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request to ping endpoint: %v", err)
		}
		defer resp.Body.Close()

		// Check status code
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}

		// Parse response
		var response struct {
			Message string `json:"message"`
			Data    struct {
				PingMessage string `json:"ping_message"`
			} `json:"data"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		// Verify response structure
		if response.Message != "success" {
			t.Errorf("Expected message 'success', got '%s'", response.Message)
		}

		if response.Data.PingMessage != "ping from repository" {
			t.Errorf("Expected ping_message 'ping from repository', got '%s'", response.Data.PingMessage)
		}

		t.Logf("Ping endpoint test passed. Response: %+v", response)
	})

	t.Run("Ping Endpoint Performance", func(t *testing.T) {
		// Test response time
		start := time.Now()
		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		duration := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Check if response time is reasonable (less than 1 second)
		if duration > time.Second {
			t.Errorf("Response time too slow: %v", duration)
		}

		t.Logf("Ping endpoint response time: %v", duration)
	})

	t.Run("Ping Endpoint Headers", func(t *testing.T) {
		// Test response headers
		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Check content type
		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			t.Error("Expected Content-Type header to be set")
		}

		t.Logf("Response headers: %+v", resp.Header)
	})
}

// TestPingHandlerErrorScenarios tests error handling scenarios
func TestPingHandlerErrorScenarios(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	baseURL := getTestBaseURL()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	t.Run("Non-existent Endpoint", func(t *testing.T) {
		// Test non-existent endpoint
		resp, err := client.Get(fmt.Sprintf("%s/non-existent", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Should return 404
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status code 404, got %d", resp.StatusCode)
		}
	})

	t.Run("Wrong HTTP Method", func(t *testing.T) {
		// Test POST to GET endpoint
		resp, err := client.Post(fmt.Sprintf("%s/ping", baseURL), "application/json", nil)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Should return 404 or 405 (Method Not Allowed)
		if resp.StatusCode != http.StatusNotFound && resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code 404 or 405, got %d", resp.StatusCode)
		}
	})
}

// TestPingHandlerLoadTesting performs basic load testing
func TestPingHandlerLoadTesting(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	t.Run("Ping Handler Load Test", func(t *testing.T) {
		baseURL := getTestBaseURL()
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		// Perform multiple concurrent requests
		const numRequests = 10
		results := make(chan bool, numRequests)

		for i := 0; i < numRequests; i++ {
			go func() {
				resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
				if err != nil {
					results <- false
					return
				}
				defer resp.Body.Close()
				results <- resp.StatusCode == http.StatusOK
			}()
		}

		// Collect results
		successCount := 0
		for i := 0; i < numRequests; i++ {
			if <-results {
				successCount++
			}
		}

		// Verify all requests succeeded
		if successCount != numRequests {
			t.Errorf("Expected all %d requests to succeed, but only %d succeeded", numRequests, successCount)
		}

		t.Logf("Load test completed: %d/%d requests succeeded", successCount, numRequests)
	})
}

// Helper function to get test base URL
func getTestBaseURL() string {
	baseURL := os.Getenv("TEST_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:1996"
	}
	return baseURL
}
