package services

// func TestApiServiceInitialize(t *testing.T) {
// 	svc := NewApiService(nil)

// 	err := svc.Initialize()
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	apiUrl := svc.ApiUrl()
// 	if apiUrl == "" {
// 		t.Fatalf("expected apiUrl to be set, got empty string")
// 	}

// 	// Optionally, validate the format of the URL
// 	if apiUrl[:10] != "localhost:" {
// 		t.Fatalf("expected apiUrl to start with 'localhost:', got %s", apiUrl)
// 	}
// }

// func TestIpfsServiceInitialize(t *testing.T) {
// 	originalIsPortAvailable := isPortAvailable
// 	isPortAvailable = func(port string) bool {
// 		return false // Simulate the port being unavailable (daemon running)
// 	}
// 	defer func() { isPortAvailable = originalIsPortAvailable }()

// 	svc := NewIpfsService(nil)

// 	err := svc.Initialize()
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if !svc.WasRunning() {
// 		t.Fatalf("expected wasRunning to be true when daemon is already running")
// 	}
// }

// func TestMonitorServiceConfiguration(t *testing.T) {
// 	chains := []string{"mainnet", "polygon"}

// 	// Test the original constructor with default values
// 	defaultSleep := 12 * time.Second
// 	defaultBatch := 200
// 	svc1 := NewMonitorService(nil, chains, defaultSleep, defaultBatch)

// 	if svc1.GetSleepTime() != defaultSleep {
// 		t.Fatalf("expected sleepTime to be %v, got %v", defaultSleep, svc1.GetSleepTime())
// 	}

// 	if svc1.GetBatchSize() != defaultBatch {
// 		t.Fatalf("expected batchSize to be %v, got %v", defaultBatch, svc1.GetBatchSize())
// 	}

// 	if len(svc1.GetChains()) != 2 {
// 		t.Fatalf("expected 2 chains, got %v", len(svc1.GetChains()))
// 	}

// 	// Test the new constructor with custom values
// 	customSleep := 30 * time.Second
// 	customBatch := 50

// 	svc2 := NewMonitorService(nil, chains, customSleep, customBatch)

// 	if svc2.GetSleepTime() != customSleep {
// 		t.Fatalf("expected custom sleepTime to be %v, got %v", customSleep, svc2.GetSleepTime())
// 	}

// 	if svc2.GetBatchSize() != customBatch {
// 		t.Fatalf("expected custom batchSize to be %v, got %v", customBatch, svc2.GetBatchSize())
// 	}

// 	if len(svc2.GetChains()) != 2 {
// 		t.Fatalf("expected 2 chains, got %v", len(svc2.GetChains()))
// 	}
// }
