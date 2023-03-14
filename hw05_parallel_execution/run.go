package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type result struct {
	numberTask int
	status     string
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	channel := make(chan result, n)
	quit := make(chan bool)

	go func() {
		defer close(channel)
		var wg sync.WaitGroup
		numGoroutines := 0
		for i := 0; i < len(tasks); i++ {
			wg.Add(1)
			go func(i int, c chan result) {
				defer func() {
					wg.Done()
					// numGoroutines--
				}()
				err := tasks[i]()
				res := result{}
				res.numberTask = i
				if err != nil {
					res.status = err.Error()
					c <- res
				} else {
					res.status = "Ok"
					c <- res
				}
			}(i, channel)
			numGoroutines++
			select {
			case <-quit:
				wg.Wait()
				return
			default:
				if numGoroutines >= n {
					wg.Wait()
					numGoroutines = 0
				}
			}
		}
	}()

	var res error
	var failed int

	for x := range channel {
		if x.status == "Ok" {
			continue
		} else {
			failed++
		}
		if failed >= m && res == nil {
			quit <- true
			res = ErrErrorsLimitExceeded
		}
	}

	// Place your code here.
	return res
}
