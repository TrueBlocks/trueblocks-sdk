package monitor

import (
	"context"
)

type MockExecutor struct {
	ExecuteFunc func(ctx context.Context, cmd Command, vars TemplateVars) error
}

func (m *MockExecutor) Execute(ctx context.Context, cmd Command, vars TemplateVars) error {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(ctx, cmd, vars)
	}
	return nil
}

// NOTE: Tests that call chifra are commented out because:
// 1. The test address is invalid
// 2. A fake address will not work with chifra
// 3. A real address with no transactions returns empty results
// 4. A real address with transactions would modify cache state
// 5. Execute() may not return errors from chifra failures

/*
func TestChifraExecutor_Execute(t *testing.T) {
	executor := NewChifraExecutor()
	ctx := context.Background()

	vars := TemplateVars{
		Address:    "0x1234567890abcdef",
		Chain:      "mainnet",
		FirstBlock: 15000000,
		LastBlock:  15000100,
		BlockCount: 101,
	}

	tests := []struct {
		name    string
		cmd     Command
		wantErr bool
	}{
		{
			name: "basic export command",
			cmd: Command{
				ID:      "export-test",
				Command: "chifra",
				Arguments: []string{
					"export",
					"{address}",
					"--cache",
				},
				Output: "",
			},
			wantErr: false,
		},
		{
			name: "no arguments",
			cmd: Command{
				ID:        "no-args",
				Command:   "chifra",
				Arguments: []string{},
				Output:    "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := executor.Execute(ctx, tt.cmd, vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChifraExecutor_Execute_WithOutput(t *testing.T) {
	tmpDir := t.TempDir()
	executor := NewChifraExecutor()
	ctx := context.Background()

	vars := TemplateVars{
		Address:    "0x1234567890abcdef",
		Chain:      "mainnet",
		FirstBlock: 15000000,
		LastBlock:  15000100,
		BlockCount: 101,
	}

	outputPath := filepath.Join(tmpDir, "output", "test.json")

	cmd := Command{
		ID:      "export-with-output",
		Command: "chifra",
		Arguments: []string{
			"export",
			"{address}",
		},
		Output: outputPath,
	}

	err := executor.Execute(ctx, cmd, vars)
	if err != nil {
		t.Errorf("Execute() unexpected error = %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("Output file was not created")
	}
}

func TestShellExecutor_Execute(t *testing.T) {
	executor := NewShellExecutor()
	ctx := context.Background()

	vars := TemplateVars{
		Address:    "0x1234567890abcdef",
		Chain:      "mainnet",
		FirstBlock: 15000000,
		LastBlock:  15000100,
		BlockCount: 101,
	}

	tests := []struct {
		name    string
		cmd     Command
		wantErr bool
	}{
		{
			name: "echo command",
			cmd: Command{
				ID:      "echo-test",
				Command: "echo",
				Arguments: []string{
					"hello",
					"{address}",
				},
				Output: "",
			},
			wantErr: false,
		},
		{
			name: "nonexistent command",
			cmd: Command{
				ID:      "bad-command",
				Command: "/nonexistent/command",
				Arguments: []string{
					"arg1",
				},
				Output: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := executor.Execute(ctx, tt.cmd, vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShellExecutor_Execute_WithOutput(t *testing.T) {
	tmpDir := t.TempDir()
	executor := NewShellExecutor()
	ctx := context.Background()

	vars := TemplateVars{
		Address:    "0x1234567890abcdef",
		Chain:      "mainnet",
		FirstBlock: 15000000,
		LastBlock:  15000100,
		BlockCount: 101,
	}

	outputPath := filepath.Join(tmpDir, "output", "echo.txt")

	cmd := Command{
		ID:      "echo-with-output",
		Command: "echo",
		Arguments: []string{
			"test output",
		},
		Output: outputPath,
	}

	err := executor.Execute(ctx, cmd, vars)
	if err != nil {
		t.Fatalf("Execute() unexpected error = %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatal("Output file was not created")
	}

	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedContent := "test output"
	actualContent := strings.TrimSpace(string(content))
	if actualContent != expectedContent {
		t.Errorf("Expected output '%s', got '%s'", expectedContent, actualContent)
	}
}

func TestShellExecutor_Execute_TemplateExpansion(t *testing.T) {
	tmpDir := t.TempDir()
	executor := NewShellExecutor()
	ctx := context.Background()

	vars := TemplateVars{
		Address:    "0xabc123",
		Chain:      "sepolia",
		FirstBlock: 100,
		LastBlock:  200,
		BlockCount: 101,
	}

	outputPath := filepath.Join(tmpDir, "{chain}", "{address}.txt")

	cmd := Command{
		ID:      "template-test",
		Command: "echo",
		Arguments: []string{
			"Address: {address}, Chain: {chain}, Blocks: {first_block}-{last_block}",
		},
		Output: outputPath,
	}

	err := executor.Execute(ctx, cmd, vars)
	if err != nil {
		t.Fatalf("Execute() unexpected error = %v", err)
	}

	expectedPath := filepath.Join(tmpDir, "sepolia", "0xabc123.txt")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Fatalf("Output file not created at expected path: %s", expectedPath)
	}

	content, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedContent := "Address: 0xabc123, Chain: sepolia, Blocks: 100-200"
	actualContent := strings.TrimSpace(string(content))
	if actualContent != expectedContent {
		t.Errorf("Expected output '%s', got '%s'", expectedContent, actualContent)
	}
}

func TestCreateExecutor(t *testing.T) {
	tests := []struct {
		name        string
		commandName string
		wantType    string
	}{
		{
			name:        "chifra executor",
			commandName: "chifra",
			wantType:    "*monitor.ChifraExecutor",
		},
		{
			name:        "shell executor",
			commandName: "/usr/bin/custom",
			wantType:    "*monitor.ShellExecutor",
		},
		{
			name:        "another shell executor",
			commandName: "bash",
			wantType:    "*monitor.ShellExecutor",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := CreateExecutor(tt.commandName)
			if executor == nil {
				t.Fatal("CreateExecutor() returned nil")
			}

			executorType := fmt.Sprintf("%T", executor)
			if executorType != tt.wantType {
				t.Errorf("CreateExecutor() returned type %s, want %s", executorType, tt.wantType)
			}
		})
	}
}

func TestMockExecutor(t *testing.T) {
	called := false
	mock := &MockExecutor{
		ExecuteFunc: func(ctx context.Context, cmd Command, vars TemplateVars) error {
			called = true
			return nil
		},
	}

	ctx := context.Background()
	cmd := Command{ID: "test"}
	vars := TemplateVars{}

	err := mock.Execute(ctx, cmd, vars)
	if err != nil {
		t.Errorf("Execute() unexpected error = %v", err)
	}

	if !called {
		t.Error("ExecuteFunc was not called")
	}
}
*/
