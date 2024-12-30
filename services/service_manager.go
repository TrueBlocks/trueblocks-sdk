package services

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ServiceManager struct {
	services       []Servicer
	stopChan       chan os.Signal
	forceExit      chan struct{}
	interruptCount int
	logger         *slog.Logger
}

func NewServiceManager(services []Servicer, logger *slog.Logger) *ServiceManager {
	return &ServiceManager{
		services:  services,
		stopChan:  make(chan os.Signal, 1),
		forceExit: make(chan struct{}),
		logger:    logger,
	}
}

func (sm *ServiceManager) StartAllServices() error {
	sm.logger.Info("Starting all services...")
	for _, svc := range sm.services {
		// Use StartService to manage each service lifecycle properly
		go func(s Servicer) {
			StartService(s, sm.stopChan)
		}(svc)
		sm.logger.Info("Service started", "name", svc.Name())
	}
	sm.logger.Info("All services started successfully")
	return nil
}

func (sm *ServiceManager) HandleSignals() {
	signal.Notify(sm.stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for sig := range sm.stopChan {
			sm.interruptCount++
			sm.logger.Info("Signal received", "signal", sig, "count", sm.interruptCount)

			if sm.interruptCount == 1 {
				sm.logger.Info("Starting cleanup for all services...")
				sm.cleanupAllServices()
				os.Exit(0)
			} else if sm.interruptCount >= 3 {
				sm.logger.Warn("Force exit after multiple interrupt signals.")
				close(sm.forceExit)
				os.Exit(1)
			}
		}
	}()
}

func (sm *ServiceManager) cleanupAllServices() {
	var wg sync.WaitGroup
	for _, svc := range sm.services {
		wg.Add(1)
		go func(service Servicer) {
			defer wg.Done()
			sm.logger.Info("Cleaning up service", "name", service.Name())
			service.Cleanup()
			sm.logger.Info("Service cleanup completed", "name", service.Name())
		}(svc)
	}
	wg.Wait()
	sm.logger.Info("All services cleaned up")
}

func (sm *ServiceManager) IsPaused(name string) ([]map[string]string, error) {
	var results []map[string]string

	for _, svc := range sm.services {
		if pauser, ok := svc.(Pauser); ok {
			if name == "" || svc.Name() == name {
				status := "running"
				if pauser.IsPaused() {
					status = "paused"
				}
				results = append(results, map[string]string{"name": svc.Name(), "status": status})
				if name != "" {
					return results, nil
				}
			}
		}
	}

	if name != "" && len(results) == 0 {
		return nil, fmt.Errorf("service '%s' not found or is not pausable", name)
	}

	return results, nil
}

func (sm *ServiceManager) Pause(name string) ([]map[string]string, error) {
	var results []map[string]string

	for _, svc := range sm.services {
		if pauser, ok := svc.(Pauser); ok {
			if name == "" || svc.Name() == name {
				if childManager, ok := svc.(ChildManager); ok && childManager.HasChild() {
					childManager.Cleanup()
					results = append(results, map[string]string{"name": svc.Name(), "status": "child process cleaned up"})
				}
				status := "not paused"
				_ = pauser.Pause()
				if pauser.IsPaused() {
					status = "paused"
				}
				results = append(results, map[string]string{"name": svc.Name(), "status": status})
				if name != "" {
					return results, nil
				}
			}
		}
	}

	if name != "" && len(results) == 0 {
		return nil, fmt.Errorf("service '%s' not found or is not pausable", name)
	}

	return results, nil
}

func (sm *ServiceManager) Unpause(name string) ([]map[string]string, error) {
	var results []map[string]string

	for _, svc := range sm.services {
		if pauser, ok := svc.(Pauser); ok {
			if name == "" || svc.Name() == name {
				if pauser.IsPaused() {
					if childManager, ok := svc.(ChildManager); ok && childManager.HasChild() {
						go func(s Servicer) {
							StartService(s, sm.stopChan)
						}(svc)
						results = append(results, map[string]string{"name": svc.Name(), "status": "child process was restarted"})
					}
					status := "not unpaused"
					_ = pauser.Unpause()
					if !pauser.IsPaused() {
						status = "unpaused"
					}
					results = append(results, map[string]string{"name": svc.Name(), "status": status})
				} else {
					results = append(results, map[string]string{"name": svc.Name(), "status": "already running"})
				}
				if name != "" {
					return results, nil
				}
			}
		}
	}

	if name != "" && len(results) == 0 {
		return nil, fmt.Errorf("service '%s' not found or is not pausable", name)
	}

	return results, nil
}

func (sm *ServiceManager) recreateService(svc Servicer) Servicer {
	return svc
}
