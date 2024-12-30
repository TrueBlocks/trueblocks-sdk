package services

import (
	"context"
	"log/slog"
	"strings"
	"time"
)

type MonitorService struct {
	paused bool
	logger *slog.Logger
	ctx    context.Context
	cancel context.CancelFunc
}

func NewMonitorService(logger *slog.Logger) *MonitorService {
	ctx, cancel := context.WithCancel(context.Background())
	return &MonitorService{
		logger: initializeLogger(logger),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *MonitorService) Name() string {
	return "monitor"
}

func (s *MonitorService) Initialize() error {
	s.logger.Info("Monitor service initialized.")
	return nil
}

func (s *MonitorService) Process(ready chan bool) error {
	s.logger.Info("Monitor service Process() invoked.")
	go func() {
		ready <- true
		for {
			select {
			case <-s.ctx.Done():
				s.logger.Info("Monitor loop stopping.")
				return
			default:
				if s.IsPaused() {
					time.Sleep(1 * time.Second)
					continue
				}
				s.logger.Info("Monitor loop running." + strings.Repeat(" ", 80))
				time.Sleep(3 * time.Second)
			}
		}
	}()
	return nil
}

func (s *MonitorService) IsPaused() bool {
	return s.paused
}

func (s *MonitorService) Pause() bool {
	s.paused = true
	return s.paused
}

func (s *MonitorService) Unpause() bool {
	s.paused = false
	return s.paused
}

func (s *MonitorService) Cleanup() {
	s.logger.Info("Monitor service cleanup started.")
	s.cancel()
	if !s.IsPaused() {
		for i := 0; i < 5; i++ {
			s.logger.Info("Monitor service cleanup in progress.", "i", i)
			time.Sleep(1 * time.Second)
		}
	}
	s.logger.Info("Monitor service cleanup complete.")
}

func (s *MonitorService) Logger() *slog.Logger {
	return s.logger
}
