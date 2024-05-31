package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, maxErrors int) error {
	if maxErrors < 0 {
		maxErrors = 0
	}
	errCount := 0
	var wg sync.WaitGroup
	taskCh := make(chan func() error)
	doneErrCh := make(chan bool, n)
	mu := sync.RWMutex{}
	isErrorsLimitExceeded := false

	worker := func() {
		defer wg.Done()
		for task := range taskCh {
			if err := task(); err != nil {
				mu.Lock()
				errCount++
				if errCount > maxErrors {
					doneErrCh <- true
					mu.Unlock()

					return
				}
				mu.Unlock()
			}
		}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker()
	}

	for _, task := range tasks {
		select {
		case taskCh <- task:
		case <-doneErrCh:
			isErrorsLimitExceeded = true
		}

		if isErrorsLimitExceeded {
			break
		}
	}
	close(taskCh)
	wg.Wait()

	if isErrorsLimitExceeded {
		return ErrErrorsLimitExceeded
	}

	return nil
}
