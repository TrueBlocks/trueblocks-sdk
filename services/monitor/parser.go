package monitor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type MonitorEntry struct {
	Address       string
	StartingBlock uint64
}

type Command struct {
	ID        string   `yaml:"id"`
	Command   string   `yaml:"command"`
	Arguments []string `yaml:"arguments"`
	Output    string   `yaml:"output"`
}

type CommandsFile struct {
	Commands []Command `yaml:"commands"`
}

func ParseWatchlist(path string) ([]MonitorEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, NewWatchlistError(path, 0, fmt.Sprintf("failed to open file: %v", err))
	}
	defer file.Close()

	var entries []MonitorEntry
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		commentIdx := strings.Index(line, "#")
		if commentIdx > 0 {
			line = strings.TrimSpace(line[:commentIdx])
		}

		parts := strings.Split(line, ",")
		address := strings.TrimSpace(parts[0])

		if !isValidAddress(address) {
			return nil, NewWatchlistError(path, lineNum, fmt.Sprintf("invalid address: %s", address))
		}

		entry := MonitorEntry{
			Address:       address,
			StartingBlock: 0,
		}

		if len(parts) > 1 {
			blockStr := strings.TrimSpace(parts[1])
			var block uint64
			_, err := fmt.Sscanf(blockStr, "%d", &block)
			if err != nil {
				return nil, NewWatchlistError(path, lineNum, fmt.Sprintf("invalid starting block: %s", blockStr))
			}
			entry.StartingBlock = block
		}

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, NewWatchlistError(path, 0, fmt.Sprintf("error reading file: %v", err))
	}

	if len(entries) == 0 {
		return nil, NewWatchlistError(path, 0, "no valid addresses found")
	}

	return entries, nil
}

func DiscoverMonitors(chain string) ([]MonitorEntry, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}

	monitorsPath := filepath.Join(homeDir, ".local", "share", "trueblocks", "cache", chain, "monitors")

	entries, err := os.ReadDir(monitorsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read monitors directory for chain %s: %v", chain, err)
	}

	var monitors []MonitorEntry
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(name, ".mon.bin") {
			continue
		}

		address := strings.TrimSuffix(name, ".mon.bin")
		if !isValidAddress(address) {
			continue
		}

		monitors = append(monitors, MonitorEntry{
			Address:       address,
			StartingBlock: 0,
		})
	}

	if len(monitors) == 0 {
		return nil, fmt.Errorf("no monitors found for chain %s", chain)
	}

	return monitors, nil
}

func ParseCommands(path string) ([]Command, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Command{}, nil
		}
		return nil, fmt.Errorf("failed to read commands file %s: %v", path, err)
	}

	var commandsFile CommandsFile
	if err := yaml.Unmarshal(data, &commandsFile); err != nil {
		return nil, fmt.Errorf("failed to parse YAML in %s: %v", path, err)
	}

	for i, cmd := range commandsFile.Commands {
		if cmd.Command == "" {
			return nil, fmt.Errorf("command at index %d missing 'command' field", i)
		}
		if len(cmd.Arguments) == 0 {
			return nil, fmt.Errorf("command '%s' has no arguments", cmd.Command)
		}
	}

	return commandsFile.Commands, nil
}

func isValidAddress(addr string) bool {
	if !strings.HasPrefix(addr, "0x") {
		return false
	}
	if len(addr) != 42 {
		return false
	}
	for _, c := range addr[2:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
