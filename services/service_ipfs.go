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

func (i *IpfsService) Name() string {
	return "IPFS Daemon"
}

func (i *IpfsService) Initialize() error {
	apiMultiaddr, err := readIPFSConfig()
	if err != nil {
		i.logger.Error("failed to read IPFS config", "error", err)
		return err
	}

	apiPort, err := extractPortFromMultiaddr(apiMultiaddr)
	if err != nil {
		i.logger.Error("failed to extract port from IPFS config", "error", err)
		return err
	}

	i.apiMultiaddr = apiMultiaddr
	i.apiPort = apiPort

	if !isPortAvailable(apiPort) {
		i.wasRunning = true
		i.logger.Info("IPFS daemon is already running", "port", apiPort)
		return nil
	}

	i.wasRunning = false
	return nil
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

func (i *IpfsService) Process(ready chan bool) error {
	if i.wasRunning {
		i.logger.Info("IPFS daemon is already running, skipping start", "multiaddr", i.apiMultiaddr)
		ready <- true
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	i.cancel = cancel

	cmd := exec.CommandContext(ctx, "ipfs", "daemon")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	cmd.Env = os.Environ()

	if err := cmd.Start(); err != nil {
		i.logger.Error("failed to start IPFS daemon", "error", err)
		cancel()
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	i.cmd = cmd
	i.logger.Info("IPFS daemon process started successfully")

	go func() {
		pollInterval := 200 * time.Millisecond
		maxRetries := 50

		for attempt := 0; attempt < maxRetries; attempt++ {
			if !isPortAvailable(i.apiPort) {
				i.logger.Info("IPFS daemon is now running", "port", i.apiPort)
				ready <- true
				return
			}
			time.Sleep(pollInterval)
		}

		i.logger.Error("timeout waiting for IPFS daemon to start", "multiaddr", i.apiMultiaddr)
		cancel()
		ready <- false
	}()

	go func() {
		if err := cmd.Wait(); err != nil {
			i.logger.Error("IPFS daemon process exited", "error", err)
		} else {
			i.logger.Info("IPFS daemon process exited cleanly")
		}
	}()

	return nil
}

func (i *IpfsService) Cleanup() {
	if i.cancel != nil {
		i.logger.Info("shutting down IPFS daemon")
		i.cancel()
	}

	if i.cmd != nil {
		if err := i.cmd.Wait(); err != nil {
			i.logger.Error("error waiting for IPFS daemon to shut down", "error", err)
		} else {
			i.logger.Info("IPFS daemon shut down cleanly")
		}
	}
}

func (i *IpfsService) Logger() *slog.Logger {
	return i.logger
}

func (i *IpfsService) ApiPort() string {
	return i.apiPort
}

func (i *IpfsService) ApiMultiaddr() string {
	return i.apiMultiaddr
}

func (i *IpfsService) WasRunning() bool {
	return i.wasRunning
}
