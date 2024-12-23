package services

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"
)

type Service interface {
	Name() string
	Initialize() error
	Process(chan bool) error
	Cleanup()
	Logger() *slog.Logger
}

func StartService(svc Service) {
	go func() {
		logger := svc.Logger()
		if logger == nil {
			logger = slog.New(slog.NewTextHandler(io.Discard, nil))
		}

		err := svc.Initialize()
		if err != nil {
			logger.Error("service initialization failed", "name", svc.Name(), "error", err)
			return
		}

		ready := make(chan bool)
		err = svc.Process(ready)
		if err != nil {
			logger.Error("service process failed", "name", svc.Name(), "error", err)
			return
		}

		if !<-ready {
			logger.Error("service did not start properly", "name", svc.Name())
			return
		}

		logger.Info("service started", "name", svc.Name())

		cleanupDone := make(chan bool, 1)
		handleSignals(svc, cleanupDone)
	}()
}

func handleSignals(svc Service, cleanupDone chan bool) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	logger := svc.Logger()
	firstSignal := true

	for {
		sig := <-sigChan
		if firstSignal {
			logger.Info("received signal, initiating cleanup", "signal", sig, "name", svc.Name())
			firstSignal = false

			go func() {
				svc.Cleanup()
				logger.Info("cleanup completed", "name", svc.Name())
				cleanupDone <- true
			}()

			<-cleanupDone
			logger.Info("exiting gracefully after cleanup", "name", svc.Name())
			os.Exit(0)
		} else {
			logger.Warn("forcing shutdown", "signal", sig, "name", svc.Name())
			os.Exit(1)
		}
	}
}

func initializeLogger(logger *slog.Logger) *slog.Logger {
	if logger == nil {
		return slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return logger
}

// getApiUrl returns the URL (including port) where the API server is running (or will run).
func getApiUrl() string {
	apiPort := strings.ReplaceAll(os.Getenv("TB_API_PORT"), ":", "")
	if apiPort == "" {
		preferred := []string{"8080", "8088", "9090", "9099"}
		apiPort = findAvailablePort(preferred)
	}

	return "localhost:" + apiPort
}

// findAvailablePort returns a port number that is available for listening.
func findAvailablePort(preferred []string) string {
	for _, port := range preferred {
		if listener, err := net.Listen("tcp", port); err == nil {
			defer listener.Close()
			return port
		}
	}

	if listener, err := net.Listen("tcp", ":0"); err == nil {
		defer listener.Close()
		addr := listener.Addr().(*net.TCPAddr)
		return fmt.Sprintf("%d", addr.Port)
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
