# gopol
A Minimal Goroutine Worker Pool in Golang

gopol

gopol is a Go library for creating and managing worker pools.
Installation

To install gopol, simply run the following command:

go

go get github.com/AlexanderMaxRanabel/gopol

Usage

go

package main

import (
	"fmt"
	"github.com/AlexanderMaxRanabel/gopol"
)

func main() {
	// Create a new worker pool with 5 workers
	pool := gopol.NewWorkerPool(5)

	// Start the worker pool
	pool.Start()
	defer pool.Stop()

	// Add 10 tasks to the worker pool
	for i := 0; i < 10; i++ {
		task := func() {
			fmt.Println("Task", i, "executed")
		}
		pool.AddTask(task)
	}
}

Contributing

If you would like to contribute to the gopol project, feel free to submit a pull request or open an issue on GitHub. All contributions are welcome!
