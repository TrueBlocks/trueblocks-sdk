package services

import (
	"io"
	"log/slog"
	"net"
	"os"
	"strings"
	"time"
)

type Servicer interface {
	Name() string
	Initialize() error
	Process(chan bool) error
	Cleanup()
	Logger() *slog.Logger
}

type Pauser interface {
	Servicer
	IsPaused() bool
	Pause() bool
	Unpause() bool
}

type ChildManager interface {
	Pauser
	HasChild() bool
	KillChild() bool
	RestartChild() bool
}

// Restarter is a marker interface for services that can be restarted
type Restarter interface {
	Servicer
}

func StartService(svc Servicer, stopChan chan os.Signal) {
	go func() {
		logger := initializeLogger(svc.Logger())
		logger.Info("Starting service", "name", svc.Name())

		if err := svc.Initialize(); err != nil {
			logger.Error("Service initialization failed", "name", svc.Name(), "error", err)
			return
		}

		ready := make(chan bool)

		go func() {
			if err := svc.Process(ready); err != nil {
				logger.Error("Service process failed", "name", svc.Name(), "error", err)
			}
		}()

		readySignal := <-ready
		if !readySignal {
			logger.Error("Service did not start properly", "name", svc.Name())
			return
		}

		logger.Info("Service started successfully", "name", svc.Name())

		cleanupDone := make(chan bool, 1)
		handleSignals(svc, stopChan, cleanupDone)
	}()
}

func handleSignals(svc Servicer, stopChan chan os.Signal, cleanupDone chan bool) {
	logger := svc.Logger()
	firstSignal := true
	for {
		sig := <-stopChan
		logger.Info("Signal received", "signal", sig, "name", svc.Name())
		if firstSignal {
			logger.Info("First signal received, initiating cleanup", "signal", sig, "name", svc.Name())
			firstSignal = false
			go func() {
				svc.Cleanup()
				logger.Info("Cleanup completed", "name", svc.Name())
				cleanupDone <- true
			}()
			<-cleanupDone
			logger.Info("Exiting gracefully after cleanup", "name", svc.Name())
			os.Exit(0)
		} else {
			logger.Warn("Additional signal received during cleanup. Ignoring.", "signal", sig, "name", svc.Name())
		}
	}
}

func initializeLogger(logger *slog.Logger) *slog.Logger {
	if logger == nil {
		return slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return logger
}

func getApiUrl() string {
	apiPort := strings.ReplaceAll(os.Getenv("TB_API_PORT"), ":", "")
	if apiPort == "" {
		apiPort = findAvailablePort([]string{"8080", "8088", "9090", "9099"})
	}
	return "localhost:" + apiPort
}

func findAvailablePort(preferred []string) string {
	for _, port := range preferred {
		if !isPortAvailable(port) {
			continue
		}
		return port
	}
	return "0"
}

var isPortAvailable = func(port string) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", port), 2*time.Second)
	if err != nil {
		return true
	}
	conn.Close()
	return false
}
