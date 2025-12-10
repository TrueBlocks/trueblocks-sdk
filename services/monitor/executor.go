package monitor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	return nil
}

type ShellExecutor struct{}

func NewShellExecutor() *ShellExecutor {
	return &ShellExecutor{}
}

func (e *ShellExecutor) Execute(ctx context.Context, cmd Command, vars TemplateVars) error {
	expandedArgs := make([]string, len(cmd.Arguments))
	for i, arg := range cmd.Arguments {
		expandedArgs[i] = ExpandTemplate(arg, vars)
	}

	output := ExpandTemplate(cmd.Output, vars)

	shellCmd := exec.CommandContext(ctx, cmd.Command, expandedArgs...)

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

		shellCmd.Stdout = file
		shellCmd.Stderr = file
	} else {
		shellCmd.Stdout = nil
		shellCmd.Stderr = nil
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
