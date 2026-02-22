package monitor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/base"
	chifraMonitor "github.com/TrueBlocks/trueblocks-chifra/v6/pkg/monitor"
	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/rpc"
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

	chain := "mainnet"
	filename := filepath.Base(path)
	if strings.HasPrefix(filename, "watchlist-") && strings.HasSuffix(filename, ".txt") {
		chain = strings.TrimSuffix(strings.TrimPrefix(filename, "watchlist-"), ".txt")
	}
	conn := rpc.TempConnection(chain)

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
		addressOrEns := strings.TrimSpace(parts[0])

		if !base.IsValidAddress(addressOrEns) {
			return nil, NewWatchlistError(path, lineNum, fmt.Sprintf("invalid address: %s", addressOrEns))
		}

		address := addressOrEns
		if strings.HasSuffix(addressOrEns, ".eth") {
			if aa, ok := conn.GetEnsAddress(addressOrEns); !ok {
				if bb, okok := conn.GetEnsAddress(addressOrEns); !okok {
					return []MonitorEntry{}, fmt.Errorf("failed to resolve ENS name: %s", addressOrEns)
				} else {
					address = bb
				}
			} else {
				address = aa
			}
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

// readMonitorState reads the LastScanned value from a monitor file using chifra's monitor package
func readMonitorState(chain, address string) (uint64, error) {
	// Create monitor instance
	mon, err := chifraMonitor.NewMonitor(chain, base.HexToAddress(address), false)
	if err != nil {
		return 0, err
	}
	defer mon.Close()

	// Read the monitor header to get LastScanned
	if err := mon.ReadMonitorHeader(); err != nil {
		// Monitor file doesn't exist yet, start from 0
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to read monitor header: %v", err)
	}

	return uint64(mon.LastScanned), nil
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

		// Read the actual LastScanned value from the monitor file
		lastScanned, err := readMonitorState(chain, address)
		if err != nil {
			// fmt.Printf("[MONITOR] Warning: Could not read state for %s: %v, using 0\n", address, err)
			lastScanned = 0
		}

		monitors = append(monitors, MonitorEntry{
			Address:       address,
			StartingBlock: lastScanned,
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
