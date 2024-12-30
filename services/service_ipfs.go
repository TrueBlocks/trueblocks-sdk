package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type IpfsService struct {
	logger       *slog.Logger
	cancel       context.CancelFunc
	cmd          *exec.Cmd
	wasRunning   bool
	apiPort      string
	apiMultiaddr string
}

func NewIpfsService(logger *slog.Logger) *IpfsService {
	return &IpfsService{
		logger: initializeLogger(logger),
	}
}

func (s *IpfsService) Name() string {
	return "ipfs"
}

func (s *IpfsService) Initialize() error {
	apiMultiaddr, err := readIPFSConfig()
	if err != nil {
		s.logger.Error("failed to read IPFS config", "error", err)
		return err
	}

	apiPort, err := extractPortFromMultiaddr(apiMultiaddr)
	if err != nil {
		s.logger.Error("failed to extract port from IPFS config", "error", err)
		return err
	}

	s.apiMultiaddr = apiMultiaddr
	s.apiPort = apiPort

	if !isPortAvailable(apiPort) {
		s.wasRunning = true
		s.logger.Info("IPFS daemon is already running", "port", apiPort)
		return nil
	}

	s.wasRunning = false
	return nil
}

func (s *IpfsService) Process(ready chan bool) error {
	if s.wasRunning {
		s.logger.Info("IPFS daemon is already running, skipping start", "multiaddr", s.apiMultiaddr)
		ready <- true
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	cmd := exec.CommandContext(ctx, "ipfs", "daemon")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	cmd.Env = os.Environ()

	if err := cmd.Start(); err != nil {
		s.logger.Error("failed to start IPFS daemon", "error", err)
		cancel()
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	s.cmd = cmd
	s.logger.Info("IPFS daemon process started successfully")

	go func() {
		pollInterval := 200 * time.Millisecond
		maxRetries := 50

		for attempt := 0; attempt < maxRetries; attempt++ {
			if !isPortAvailable(s.apiPort) {
				s.logger.Info("IPFS daemon is now running", "port", s.apiPort)
				ready <- true
				return
			}
			time.Sleep(pollInterval)
		}

		s.logger.Error("timeout waiting for IPFS daemon to start", "multiaddr", s.apiMultiaddr)
		cancel()
		ready <- false
	}()

	go func() {
		if err := cmd.Wait(); err != nil {
			s.logger.Error("IPFS daemon process exited", "error", err)
		} else {
			s.logger.Info("IPFS daemon process exited cleanly")
		}
	}()

	return nil
}

func (s *IpfsService) Cleanup() {
	if s.cancel != nil {
		s.logger.Info("shutting down IPFS daemon")
		s.cancel()
	}

	if s.cmd != nil {
		if err := s.cmd.Wait(); err != nil {
			s.logger.Error("error waiting for IPFS daemon to shut down", "error", err)
		} else {
			s.logger.Info("IPFS daemon shut down cleanly")
		}
	}
}

func (s *IpfsService) Logger() *slog.Logger {
	return s.logger
}

func (s *IpfsService) ApiPort() string {
	return s.apiPort
}

func (s *IpfsService) ApiMultiaddr() string {
	return s.apiMultiaddr
}

func (s *IpfsService) WasRunning() bool {
	return s.wasRunning
}

func readIPFSConfig() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ".ipfs", "config")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read IPFS config file: %w", err)
	}

	var config struct {
		Addresses struct {
			API string `json:"API"`
		} `json:"Addresses"`
	}

	if err := json.Unmarshal(configData, &config); err != nil {
		return "", fmt.Errorf("failed to parse IPFS config file: %w", err)
	}

	return config.Addresses.API, nil
}

func extractPortFromMultiaddr(multiaddr string) (string, error) {
	parts := strings.Split(multiaddr, "/")
	for i, part := range parts {
		if part == "tcp" && i+1 < len(parts) {
			return parts[i+1], nil
		}
	}
	return "", fmt.Errorf("no TCP port found in multiaddress: %s", multiaddr)
}
