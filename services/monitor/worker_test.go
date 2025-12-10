package monitor

import (
	"context"
	"testing"
	"time"
)

func TestNewWorkerPool(t *testing.T) {
	executor := NewShellExecutor()
	pool := NewWorkerPool(5, executor)

	if pool == nil {
		t.Fatal("NewWorkerPool returned nil")
	}

	if pool.concurrency != 5 {
		t.Errorf("Expected concurrency 5, got %d", pool.concurrency)
	}

	if pool.jobs == nil {
		t.Error("jobs channel not initialized")
	}

	if pool.results == nil {
		t.Error("results channel not initialized")
	}
}

func TestWorkerPool_StartAndStop(t *testing.T) {
	executor := NewShellExecutor()
	pool := NewWorkerPool(3, executor)

	pool.Start()

	time.Sleep(10 * time.Millisecond)

	pool.Stop()
}

func TestWorkerPool_SubmitAndWait(t *testing.T) {
	mock := &MockExecutor{
		ExecuteFunc: func(ctx context.Context, cmd Command, vars TemplateVars) error {
			return nil
		},
	}

	pool := NewWorkerPool(2, mock)
	pool.Start()

	job := MonitorJob{
		Entry: MonitorEntry{
			Address:       "0x1234567890123456789012345678901234567890",
			StartingBlock: 0,
		},
		Commands: []Command{
			{
				ID:      "test-cmd",
				Command: "echo",
				Arguments: []string{
					"test",
				},
				Output: "",
			},
		},
		Vars: TemplateVars{
			Address:    "0x1234567890123456789012345678901234567890",
			Chain:      "mainnet",
			FirstBlock: 0,
			LastBlock:  100,
			BlockCount: 101,
		},
	}

	pool.Submit(job)

	results := pool.Wait()

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	if !results[0].Success {
		t.Errorf("Expected success, got failure: %v", results[0].Error)
	}

	if results[0].Address != job.Entry.Address {
		t.Errorf("Expected address %s, got %s", job.Entry.Address, results[0].Address)
	}
}

func TestWorkerPool_MultipleJobs(t *testing.T) {
	mock := &MockExecutor{
		ExecuteFunc: func(ctx context.Context, cmd Command, vars TemplateVars) error {
			return nil
		},
	}

	pool := NewWorkerPool(3, mock)
	pool.Start()

	jobCount := 10
	for i := 0; i < jobCount; i++ {
		job := MonitorJob{
			Entry: MonitorEntry{
				Address:       "0x1234567890123456789012345678901234567890",
				StartingBlock: 0,
			},
			Commands: []Command{
				{
					ID:      "test-cmd",
					Command: "echo",
					Arguments: []string{
						"test",
					},
					Output: "",
				},
			},
			Vars: TemplateVars{
				Address:    "0x1234567890123456789012345678901234567890",
				Chain:      "mainnet",
				FirstBlock: 0,
				LastBlock:  100,
				BlockCount: 101,
			},
		}
		pool.Submit(job)
	}

	results := pool.Wait()

	if len(results) != jobCount {
		t.Errorf("Expected %d results, got %d", jobCount, len(results))
	}

	successCount := 0
	for _, result := range results {
		if result.Success {
			successCount++
		}
	}

	if successCount != jobCount {
		t.Errorf("Expected %d successful jobs, got %d", jobCount, successCount)
	}
}

func TestWorkerPool_CommandFailure(t *testing.T) {
	mock := &MockExecutor{
		ExecuteFunc: func(ctx context.Context, cmd Command, vars TemplateVars) error {
			return NewCommandError(cmd.ID, vars.Address, "simulated failure", nil)
		},
	}

	pool := NewWorkerPool(2, mock)
	pool.Start()

	job := MonitorJob{
		Entry: MonitorEntry{
			Address:       "0x1234567890123456789012345678901234567890",
			StartingBlock: 0,
		},
		Commands: []Command{
			{
				ID:      "failing-cmd",
				Command: "echo",
				Arguments: []string{
					"test",
				},
				Output: "",
			},
		},
		Vars: TemplateVars{
			Address:    "0x1234567890123456789012345678901234567890",
			Chain:      "mainnet",
			FirstBlock: 0,
			LastBlock:  100,
			BlockCount: 101,
		},
	}

	pool.Submit(job)

	results := pool.Wait()

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	if results[0].Success {
		t.Error("Expected failure, got success")
	}

	if results[0].Error == nil {
		t.Error("Expected error, got nil")
	}
}

func TestWorkerPool_MultipleCommands(t *testing.T) {
	executionCount := 0
	mock := &MockExecutor{
		ExecuteFunc: func(ctx context.Context, cmd Command, vars TemplateVars) error {
			executionCount++
			return nil
		},
	}

	pool := NewWorkerPool(2, mock)
	pool.Start()

	job := MonitorJob{
		Entry: MonitorEntry{
			Address:       "0x1234567890123456789012345678901234567890",
			StartingBlock: 0,
		},
		Commands: []Command{
			{
				ID:        "cmd1",
				Command:   "echo",
				Arguments: []string{"test1"},
				Output:    "",
			},
			{
				ID:        "cmd2",
				Command:   "echo",
				Arguments: []string{"test2"},
				Output:    "",
			},
			{
				ID:        "cmd3",
				Command:   "echo",
				Arguments: []string{"test3"},
				Output:    "",
			},
		},
		Vars: TemplateVars{
			Address:    "0x1234567890123456789012345678901234567890",
			Chain:      "mainnet",
			FirstBlock: 0,
			LastBlock:  100,
			BlockCount: 101,
		},
	}

	pool.Submit(job)

	results := pool.Wait()

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	if !results[0].Success {
		t.Errorf("Expected success, got failure: %v", results[0].Error)
	}

	if executionCount != 3 {
		t.Errorf("Expected 3 command executions, got %d", executionCount)
	}
}

func TestWorkerPool_CommandFailureStopsSequence(t *testing.T) {
	executionCount := 0
	mock := &MockExecutor{
		ExecuteFunc: func(ctx context.Context, cmd Command, vars TemplateVars) error {
			executionCount++
			if cmd.ID == "cmd2" {
				return NewCommandError(cmd.ID, vars.Address, "simulated failure", nil)
			}
			return nil
		},
	}

	pool := NewWorkerPool(2, mock)
	pool.Start()

	job := MonitorJob{
		Entry: MonitorEntry{
			Address:       "0x1234567890123456789012345678901234567890",
			StartingBlock: 0,
		},
		Commands: []Command{
			{
				ID:        "cmd1",
				Command:   "echo",
				Arguments: []string{"test1"},
				Output:    "",
			},
			{
				ID:        "cmd2",
				Command:   "echo",
				Arguments: []string{"test2"},
				Output:    "",
			},
			{
				ID:        "cmd3",
				Command:   "echo",
				Arguments: []string{"test3"},
				Output:    "",
			},
		},
		Vars: TemplateVars{
			Address:    "0x1234567890123456789012345678901234567890",
			Chain:      "mainnet",
			FirstBlock: 0,
			LastBlock:  100,
			BlockCount: 101,
		},
	}

	pool.Submit(job)

	results := pool.Wait()

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	if results[0].Success {
		t.Error("Expected failure, got success")
	}

	if executionCount != 2 {
		t.Errorf("Expected 2 command executions (stopped at failure), got %d", executionCount)
	}
}
