package monitor

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-chifra/v6/pkg/colors"
)

type CommandExecutor interface {
	Execute(ctx context.Context, cmd Command, vars TemplateVars) error
}

type ChifraExecutor struct{}

func NewChifraExecutor() *ChifraExecutor {
	return &ChifraExecutor{}
}

func (e *ChifraExecutor) Execute(ctx context.Context, cmd Command, vars TemplateVars) error {
	expandedArgs := make([]string, len(cmd.Arguments))
	for i, arg := range cmd.Arguments {
		expandedArgs[i] = ExpandTemplate(arg, vars)
	}

	output := ExpandTemplate(cmd.Output, vars)

	if len(expandedArgs) == 0 {
		return NewCommandError(cmd.ID, vars.Address, "no arguments provided", nil)
	}

	chifraCmd := expandedArgs[0]
	chifraArgs := expandedArgs[1:]

	fmt.Println(colors.BrightGreen, "output", output, "cmd", strings.Join(expandedArgs, " "), colors.Off)
	if output != "" {
		dir := filepath.Dir(output)
		if dir != "" && dir != "." {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return NewCommandError(cmd.ID, vars.Address, fmt.Sprintf("failed to create output directory: %s", dir), err)
			}
		}

		file, err := os.Create(output)
		if err != nil {
			return NewCommandError(cmd.ID, vars.Address, fmt.Sprintf("failed to create output file: %s", output), err)
		}
		defer file.Close()
	}

	cmdStr := fmt.Sprintf("chifra %s %s", chifraCmd, strings.Join(chifraArgs, " "))
	_ = cmdStr
	fmt.Println(colors.BrightGreen, cmdStr, colors.Off)

	return nil
}

type ShellExecutor struct{}

func NewShellExecutor() *ShellExecutor {
	return &ShellExecutor{}
}

func (e *ShellExecutor) Execute(ctx context.Context, cmd Command, vars TemplateVars) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		if duration > 30*time.Second {
			fmt.Printf("[MONITOR] SLOW: Command %s for %s took %v\n", cmd.ID, vars.Address, duration)
		}
	}()

	expandedArgs := make([]string, len(cmd.Arguments))
	for i, arg := range cmd.Arguments {
		expandedArgs[i] = ExpandTemplate(arg, vars)
	}

	output := ExpandTemplate(cmd.Output, vars)
	fmt.Println(colors.BrightGreen, cmd, strings.Join(expandedArgs, " "), colors.Off)

	// Create timeout context for this specific command (5 minutes max per monitor)
	cmdCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	shellCmd := exec.CommandContext(cmdCtx, cmd.Command, expandedArgs...)

	if output != "" && output != "/dev/null" {
		dir := filepath.Dir(output)
		if dir != "" && dir != "." {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return NewCommandError(cmd.ID, vars.Address, fmt.Sprintf("failed to create output directory: %s", dir), err)
			}
		}

		file, err := os.Create(output)
		if err != nil {
			return NewCommandError(cmd.ID, vars.Address, fmt.Sprintf("failed to create output file: %s", output), err)
		}
		defer file.Close()

		shellCmd.Stdout = file
		shellCmd.Stderr = file
	} else {
		// Discard all output - don't set to nil as that can cause blocking
		shellCmd.Stdout = io.Discard
		shellCmd.Stderr = io.Discard
	}

	if err := shellCmd.Run(); err != nil {
		return NewCommandError(cmd.ID, vars.Address, "command execution failed", err)
	}

	return nil
}

func CreateExecutor(commandName string) CommandExecutor {
	if commandName == "chifra" {
		return NewChifraExecutor()
	}
	return NewShellExecutor()
}
