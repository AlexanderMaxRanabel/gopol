package main

import (
	"fmt"
	"sync"
)

// Task is the type of function that can be executed by a worker
type Task func()

// Worker is a type that executes tasks
type Worker struct {
	tasks chan Task
	wg    sync.WaitGroup
}

// NewWorker creates a new worker and returns it
func NewWorker() *Worker {
	return &Worker{
		tasks: make(chan Task),
	}
}

// Start starts the worker
func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			task, ok := <-w.tasks
			if !ok {
				return
			}
			task()
		}
	}()
}

// Stop stops the worker
func (w *Worker) Stop() {
	close(w.tasks)
	w.wg.Wait()
}

// WorkerPool is a type that manages a pool of workers
type WorkerPool struct {
	workers []*Worker
	tasks   chan Task
	wg      sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers and returns it
func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		workers: make([]*Worker, numWorkers),
		tasks:   make(chan Task),
	}
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker()
		pool.workers[i] = worker
		worker.Start()
	}
	return pool
}

// Start starts the worker pool
func (p *WorkerPool) Start() {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		for {
			task, ok := <-p.tasks
			if !ok {
				return
			}
			worker := p.workers[0]
			worker.tasks <- task
			p.workers = append(p.workers[1:], worker)
		}
	}()
}

// Stop stops the worker pool
func (p *WorkerPool) Stop() {
	close(p.tasks)
	p.wg.Wait()
	for _, worker := range p.workers {
		worker.Stop()
	}
}

// AddTask adds a task to the worker pool
func (p *WorkerPool) AddTask(task Task) {
	p.tasks <- task
}

func main() {
	pool := NewWorkerPool(5)
	pool.Start()
	defer pool.Stop()

	for i := 0; i < 10; i++ {
		task := func() {
			fmt.Println("Task", i, "executed")
		}
		pool.AddTask(task)
	}
}

