//go:build integration
// +build integration

package ping_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestPingModuleIntegration tests the complete ping module end-to-end
func TestPingModuleIntegration(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	// Configuration
	baseURL := getTestBaseURL()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	t.Run("Complete Ping Module Flow", func(t *testing.T) {
		// Test the complete flow: HTTP Request -> Handler -> Service -> Repository -> Response

		// Make HTTP request
		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request to ping endpoint: %v", err)
		}
		defer resp.Body.Close()

		// Verify HTTP response
		assert.Equal(t, http.StatusOK, resp.StatusCode)

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
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "ping from repository", response.Data.PingMessage)

		t.Logf("Complete ping module flow test passed. Response: %+v", response)
	})

	t.Run("Ping Module Performance", func(t *testing.T) {
		// Test the performance of the complete module

		start := time.Now()
		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		duration := time.Since(start)

		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Verify performance
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Less(t, duration, time.Second, "Complete module should respond within 1 second")

		t.Logf("Ping module response time: %v", duration)
	})

	t.Run("Ping Module Error Handling", func(t *testing.T) {
		// Test error handling scenarios

		// Test with invalid endpoint
		resp, err := client.Get(fmt.Sprintf("%s/ping/invalid", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Should return 404
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Ping Module Headers and Metadata", func(t *testing.T) {
		// Test response headers and metadata

		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// Check important headers
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		contentType := resp.Header.Get("Content-Type")
		assert.NotEmpty(t, contentType, "Content-Type header should be set")

		t.Logf("Response headers: %+v", resp.Header)
	})
}

// TestPingModuleWithDatabase tests the module with real database connectivity
func TestPingModuleWithDatabase(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	t.Run("Ping Module with Database Connectivity", func(t *testing.T) {
		// This test would verify that the module can connect to the database
		// For now, we'll test the basic functionality

		baseURL := getTestBaseURL()
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// The module should work even without database connectivity
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// TODO: Add database connectivity verification
		// 1. Check if database is accessible
		// 2. Verify database ping functionality
		// 3. Test database error scenarios
	})
}

// TestPingModuleWithCache tests the module with real cache connectivity
func TestPingModuleWithCache(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	t.Run("Ping Module with Cache Connectivity", func(t *testing.T) {
		// This test would verify that the module can connect to the cache
		// For now, we'll test the basic functionality

		baseURL := getTestBaseURL()
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := client.Get(fmt.Sprintf("%s/ping", baseURL))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		// The module should work even without cache connectivity
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// TODO: Add cache connectivity verification
		// 1. Check if cache is accessible
		// 2. Verify cache ping functionality
		// 3. Test cache error scenarios
	})
}

// TestPingModuleLoadTesting performs basic load testing
func TestPingModuleLoadTesting(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST=true to run")
	}

	t.Run("Ping Module Load Test", func(t *testing.T) {
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
		assert.Equal(t, numRequests, successCount, "All concurrent requests should succeed")

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
