package services

import (
	"testing"
)

func TestApiServiceInitialize(t *testing.T) {
	svc := NewApiService(nil)

	err := svc.Initialize()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	apiUrl := svc.ApiUrl()
	if apiUrl == "" {
		t.Fatalf("expected apiUrl to be set, got empty string")
	}

	// Optionally, validate the format of the URL
	if apiUrl[:10] != "localhost:" {
		t.Fatalf("expected apiUrl to start with 'localhost:', got %s", apiUrl)
	}
}

func TestIpfsServiceInitialize(t *testing.T) {
	originalIsPortAvailable := isPortAvailable
	isPortAvailable = func(port string) bool {
		return false // Simulate the port being unavailable (daemon running)
	}
	defer func() { isPortAvailable = originalIsPortAvailable }()

	svc := NewIpfsService(nil)

	err := svc.Initialize()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !svc.WasRunning() {
		t.Fatalf("expected wasRunning to be true when daemon is already running")
	}
}
