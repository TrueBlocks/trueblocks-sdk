package monitor

import (
	"fmt"
	"runtime"
)

type MonitorConfig struct {
	WatchlistDir    string
	CommandsDir     string
	BatchSize       int
	Concurrency     int
	Sleep           int
	MaxBlocksPerRun int
}

func NewMonitorConfig() *MonitorConfig {
	return &MonitorConfig{
		BatchSize:       8,
		Concurrency:     max(1, runtime.NumCPU()/2),
		Sleep:           12,
		MaxBlocksPerRun: 0,
	}
}

func (c *MonitorConfig) Validate() error {
	if c.WatchlistDir == "" {
		return fmt.Errorf("WatchlistDir cannot be empty")
	}
	if c.CommandsDir == "" {
		return fmt.Errorf("CommandsDir cannot be empty")
	}
	if c.BatchSize < 1 {
		return fmt.Errorf("BatchSize must be at least 1, got %d", c.BatchSize)
	}
	if c.Concurrency < 1 {
		return fmt.Errorf("concurrency must be at least 1, got %d", c.Concurrency)
	}
	if c.Sleep < 0 {
		return fmt.Errorf("sleep must be non-negative, got %d", c.Sleep)
	}
	if c.MaxBlocksPerRun < 0 {
		return fmt.Errorf("MaxBlocksPerRun must be non-negative, got %d", c.MaxBlocksPerRun)
	}
	return nil
}
