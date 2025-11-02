package services

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/base"
)

type ControlService struct {
	logger      *slog.Logger
	manager     *ServiceManager
	server      *http.Server
	listenAddr  string
	port        string
	ctx         context.Context
	cancel      context.CancelFunc
	extra       []func(*http.ServeMux)
	rootHandler http.HandlerFunc
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

func (s *ControlService) Port() int {
	return int(base.MustParseInt64(s.port))
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
	mux.HandleFunc("/restart", s.handleRestart)
	if s.rootHandler == nil {
		s.rootHandler = s.handleDefault
	}
	mux.HandleFunc("/", s.rootHandler)
	for _, add := range s.extra {
		add(mux)
	}
	s.server = &http.Server{Addr: s.listenAddr, Handler: mux}
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
	s.ctx, s.cancel = context.WithCancel(context.Background())
	if s.server != nil {
		_ = s.server.Close()
	}
}

func (s *ControlService) Logger() *slog.Logger {
	return s.logger
}

func (s *ControlService) handleDefault(w http.ResponseWriter, r *http.Request) {
	_ = r
	results := map[string]string{
		"/status":   "[name]",
		"/isPaused": "name",
		"/pause":    "name",
		"/unpause":  "name",
		"/restart":  "name",
	}
	writeJSONResponse(w, results)
}

// Extension points
func (s *ControlService) AddHandler(pattern string, h http.HandlerFunc) {
	s.extra = append(s.extra, func(mux *http.ServeMux) { mux.HandleFunc(pattern, h) })
}

func (s *ControlService) SetRootHandler(h http.HandlerFunc) { s.rootHandler = h }

func (s *ControlService) DefaultRootHandler() http.HandlerFunc { return s.handleDefault }

func (s *ControlService) handleIsPaused(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" || name == "all" {
		s.logger.Info("Received status request for all services", "remote_addr", r.RemoteAddr)
	} else {
		s.logger.Info("Received status request", "service", name, "remote_addr", r.RemoteAddr)
	}

	results, err := s.manager.IsPaused(name)
	if err != nil {
		s.logger.Error("Status request failed", "service", name, "error", err.Error())
		writeJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, result := range results {
		serviceName := result["name"]
		status := result["status"]
		s.logger.Info("Service status result", "service", serviceName, "status", status)
	}

	writeJSONResponse(w, results)
}

func (s *ControlService) handlePause(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" || name == "all" {
		s.logger.Info("Received pause request for all services", "remote_addr", r.RemoteAddr)
	} else {
		s.logger.Info("Received pause request", "service", name, "remote_addr", r.RemoteAddr)
	}

	results, err := s.manager.Pause(name)
	if err != nil {
		s.logger.Error("Pause request failed", "service", name, "error", err.Error())
		writeJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, result := range results {
		serviceName := result["name"]
		status := result["status"]
		s.logger.Info("Service paused", "service", serviceName, "status", status)
	}

	writeJSONResponse(w, results)
}

func (s *ControlService) handleUnpause(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" || name == "all" {
		s.logger.Info("Received unpause request for all services", "remote_addr", r.RemoteAddr)
	} else {
		s.logger.Info("Received unpause request", "service", name, "remote_addr", r.RemoteAddr)
	}

	results, err := s.manager.Unpause(name)
	if err != nil {
		s.logger.Error("Unpause request failed", "service", name, "error", err.Error())
		writeJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, result := range results {
		serviceName := result["name"]
		status := result["status"]
		s.logger.Info("Service unpaused", "service", serviceName, "status", status)
	}

	writeJSONResponse(w, results)
}

func (s *ControlService) handleRestart(w http.ResponseWriter, r *http.Request) {
	if s.manager == nil {
		s.logger.Error("Service manager not attached")
		writeJSONErrorResponse(w, "Service manager not attached", http.StatusInternalServerError)
		return
	}

	serviceName := r.URL.Query().Get("name")
	if serviceName == "" {
		serviceName = "all"
	}

	s.logger.Info("Received restart request", "remote_addr", r.RemoteAddr, "service", serviceName)

	results, err := s.manager.Restart(serviceName)
	if err != nil {
		s.logger.Error("Restart request failed", "service", serviceName, "error", err.Error())
		writeJSONErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, result := range results {
		svcName := result["name"]
		status := result["status"]
		s.logger.Info("Service restart result", "service", svcName, "status", status)
	}

	writeJSONResponse(w, results)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func writeJSONErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
