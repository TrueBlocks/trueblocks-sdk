package services

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ControlService struct {
	logger     *slog.Logger
	manager    *ServiceManager
	server     *http.Server
	listenAddr string
	port       string
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewControlService(logger *slog.Logger) *ControlService {
	port := findAvailablePort([]string{"8338", "8337", "8336", "8335"})
	ctx, cancel := context.WithCancel(context.Background())
	return &ControlService{
		manager:    nil,
		logger:     logger,
		port:       port,
		listenAddr: ":" + port,
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (s *ControlService) Name() string {
	return "control"
}

func (s *ControlService) AttachServiceManager(manager *ServiceManager) {
	s.manager = manager
}

func (s *ControlService) Initialize() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", s.handleIsPaused)
	mux.HandleFunc("/isPaused", s.handleIsPaused)
	mux.HandleFunc("/pause", s.handlePause)
	mux.HandleFunc("/unpause", s.handleUnpause)
	mux.HandleFunc("/", s.handleDefault)

	s.server = &http.Server{
		Addr:    s.listenAddr,
		Handler: mux,
	}

	return nil
}

func (s *ControlService) Process(ready chan bool) error {
	ready <- true

	s.logger.Info("Control Service starting", "address", s.listenAddr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *ControlService) Cleanup() {
	s.cancel()
	if s.server != nil {
		s.server.Close()
	}
}

func (s *ControlService) Logger() *slog.Logger {
	return s.logger
}

func (s *ControlService) handleDefault(w http.ResponseWriter, r *http.Request) {
	results := map[string]string{
		"/status":   "[name]",
		"/isPaused": "name",
		"/pause":    "name",
		"/unpause":  "name",
	}
	writeJSONResponse(w, results)
}

func (s *ControlService) handleIsPaused(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	results, err := s.manager.IsPaused(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSONResponse(w, results)
}

func (s *ControlService) handlePause(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	results, err := s.manager.Pause(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSONResponse(w, results)
}

func (s *ControlService) handleUnpause(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	results, err := s.manager.Unpause(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSONResponse(w, results)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
