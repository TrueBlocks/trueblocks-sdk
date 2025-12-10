package monitor

import "fmt"

type ConfigError struct {
	Field   string
	Message string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("config error [%s]: %s", e.Field, e.Message)
}

func NewConfigError(field, message string) error {
	return &ConfigError{
		Field:   field,
		Message: message,
	}
}

type WatchlistError struct {
	Path    string
	Line    int
	Message string
}

func (e *WatchlistError) Error() string {
	if e.Line > 0 {
		return fmt.Sprintf("watchlist error [%s:%d]: %s", e.Path, e.Line, e.Message)
	}
	return fmt.Sprintf("watchlist error [%s]: %s", e.Path, e.Message)
}

func NewWatchlistError(path string, line int, message string) error {
	return &WatchlistError{
		Path:    path,
		Line:    line,
		Message: message,
	}
}

type CommandError struct {
	CommandID string
	Address   string
	Message   string
	Cause     error
}

func (e *CommandError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("command error [%s] for address %s: %s: %v", e.CommandID, e.Address, e.Message, e.Cause)
	}
	return fmt.Sprintf("command error [%s] for address %s: %s", e.CommandID, e.Address, e.Message)
}

func (e *CommandError) Unwrap() error {
	return e.Cause
}

func NewCommandError(commandID, address, message string, cause error) error {
	return &CommandError{
		CommandID: commandID,
		Address:   address,
		Message:   message,
		Cause:     cause,
	}
}
