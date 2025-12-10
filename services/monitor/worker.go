package monitor

import (
	"context"
	"sync"
)

type MonitorJob struct {
	Entry    MonitorEntry
	Commands []Command
	Vars     TemplateVars
}

type JobResult struct {
	Address string
	Success bool
	Error   error
}

type WorkerPool struct {
	concurrency int
	jobs        chan MonitorJob
	results     chan JobResult
	wg          sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
	executor    CommandExecutor
}

func NewWorkerPool(concurrency int, executor CommandExecutor) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		concurrency: concurrency,
		jobs:        make(chan MonitorJob, concurrency*2),
		results:     make(chan JobResult, concurrency*2),
		ctx:         ctx,
		cancel:      cancel,
		executor:    executor,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.concurrency; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()

	for {
		select {
		case <-wp.ctx.Done():
			return
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}
			wp.processJob(job)
		}
	}
}

func (wp *WorkerPool) processJob(job MonitorJob) {
	result := JobResult{
		Address: job.Entry.Address,
		Success: true,
	}

	for _, cmd := range job.Commands {
		if err := wp.executor.Execute(wp.ctx, cmd, job.Vars); err != nil {
			result.Success = false
			result.Error = err
			break
		}
	}

	select {
	case wp.results <- result:
	case <-wp.ctx.Done():
		return
	}
}

func (wp *WorkerPool) Submit(job MonitorJob) {
	select {
	case wp.jobs <- job:
	case <-wp.ctx.Done():
		return
	}
}

func (wp *WorkerPool) Wait() []JobResult {
	close(wp.jobs)

	// Start draining results in background while workers finish
	resultsChan := make(chan []JobResult, 1)
	go func() {
		var results []JobResult
		for result := range wp.results {
			results = append(results, result)
		}
		resultsChan <- results
	}()

	// Wait for all workers to finish
	wp.wg.Wait()
	close(wp.results)

	// Get collected results
	return <-resultsChan
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.cancel()
	wp.wg.Wait()
}
