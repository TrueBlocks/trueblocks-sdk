// Package services provides a framework for managing long-running background services
// with lifecycle management, pause/unpause functionality, and restart capabilities.
//
// # Basic Usage
//
// Create and start services using the ServiceManager:
//
//	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
//
//	// Create services
//	scraper := services.NewScrapeService(logger, "blooms", []string{"mainnet"}, 13, 2000)
//	monitor := services.NewMonitorService(logger)
//	api := services.NewApiService(logger)
//	control := services.NewControlService(logger)
//
//	// Create service manager
//	serviceList := []services.Servicer{scraper, monitor, api, control}
//	manager := services.NewServiceManager(serviceList, logger)
//
//	// Attach manager to control service for HTTP API
//	control.AttachServiceManager(manager)
//
//	// Start all services
//	manager.StartAllServices()
//
//	// Handle shutdown signals
//	manager.HandleSignals()
//
// # HTTP Control API
//
// The control service provides REST endpoints for managing services:
//
//	GET /status                     - Check service status
//	POST /pause?name=service_name   - Pause a service
//	POST /unpause?name=service_name - Unpause a service
//	POST /restart?name=service_name - Restart a service
//
// Use name=all to operate on all applicable services.
//
// # Service Interfaces
//
// Services implement different interfaces based on their capabilities:
//
//	Servicer  - Basic service interface (all services)
//	Pauser    - Services that can be paused/unpaused (scraper, monitor)
//	Restarter - Services that can be restarted (scraper, monitor)
//
// # Creating Custom Services
//
// Implement the Servicer interface to create custom services:
//
//	type MyService struct {
//		logger *slog.Logger
//		ctx    context.Context
//		cancel context.CancelFunc
//	}
//
//	func (s *MyService) Name() string { return "my-service" }
//
//	func (s *MyService) Initialize() error {
//		// Setup logic here
//		return nil
//	}
//
//	func (s *MyService) Process(ready chan bool) error {
//		ready <- true
//		for {
//			select {
//			case <-s.ctx.Done():
//				return nil
//			default:
//				// Main service logic here
//			}
//		}
//	}
//
//	func (s *MyService) Cleanup() {
//		s.cancel()
//		// Reset context for potential restart
//		s.ctx, s.cancel = context.WithCancel(context.Background())
//	}
//
//	func (s *MyService) Logger() *slog.Logger { return s.logger }
//
// For more information about the TrueBlocks SDK, see the parent README.
package services
