package test

import (
    "net/http"
    "testing"
)

func TestHealthEndpoint(t *testing.T) {
    resp, err := http.Get("http://localhost:8080/health")
    if err != nil || resp.StatusCode != http.StatusOK {
        t.Errorf("Health endpoint failed: %v", err)
    }
}
