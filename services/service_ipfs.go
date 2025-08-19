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
	ctx          context.Context
	cancel       context.CancelFunc
	cmd          *exec.Cmd
	wasRunning   bool
	apiPort      string
	apiMultiaddr string
	processDone  chan struct{}
}

func NewIpfsService(logger *slog.Logger) *IpfsService {
	ctx, cancel := context.WithCancel(context.Background())
	return &IpfsService{
		logger:      initializeLogger(logger),
		ctx:         ctx,
		cancel:      cancel,
		processDone: make(chan struct{}),
	}
}

func (s *IpfsService) Name() string {
	return "ipfs"
}

func (s *IpfsService) Initialize() error {
	s.logger.Info("Initializing IPFS service...")
	apiMultiaddr, err := readIPFSConfig()
	if err != nil {
		s.logger.Error("Failed to read IPFS config", "error", err)
		return err
	}

	apiPort, err := extractPortFromMultiaddr(apiMultiaddr)
	if err != nil {
		s.logger.Error("Failed to extract port from IPFS config", "error", err)
		return err
	}

	s.apiMultiaddr = apiMultiaddr
	s.apiPort = apiPort

	if !isPortAvailable(apiPort) {
		s.wasRunning = true
		s.logger.Info("IPFS daemon is already running", "port", apiPort, "multiaddr", apiMultiaddr)
		return nil
	}

	s.wasRunning = false
	s.logger.Info("IPFS service initialized", "port", apiPort, "multiaddr", apiMultiaddr, "wasRunning", s.wasRunning)
	return nil
}

func (s *IpfsService) Process(ready chan bool) error {
	if s.wasRunning {
		ready <- true
		return nil
	}

	cmd := exec.CommandContext(s.ctx, "ipfs", "daemon")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	cmd.Env = os.Environ()

	if err := cmd.Start(); err != nil {
		s.cancel()
		return err
	}

	s.cmd = cmd

	go func() {
		defer close(s.processDone)
		_ = cmd.Wait()
	}()

	go func() {
		pollInterval := 200 * time.Millisecond
		maxRetries := 50
		for attempt := 0; attempt < maxRetries; attempt++ {
			if !isPortAvailable(s.apiPort) {
				ready <- true
				return
			}
			time.Sleep(pollInterval)
		}
		s.cancel()
		ready <- false
	}()

	return nil
}

func (s *IpfsService) Cleanup() {
	s.cancel()
	if s.cmd != nil && s.cmd.Process != nil {
		select {
		case <-s.processDone:
		case <-time.After(5 * time.Second):
			_ = s.cmd.Process.Kill()
		}
		s.cmd = nil
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
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
