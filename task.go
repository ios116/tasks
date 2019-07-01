package tasks

import (
	"log"
	"sync"
)

type function func() error

// ErrorCounter - error counter returning from functions
type ErrorCounter struct {
	sync.Mutex
	count int
}

// Add - safely adds one
func (ec *ErrorCounter) Add() int {
	ec.Lock()
	ec.count++
	ec.Unlock()
	return ec.count
}

func cancelled() bool {
	select {
	case <-stop:
		return true
	default:
		return false
	}
}

func maker(task function, n int) {
    if err := task(); err != nil {
		count := counter.Add()
		if count >= n {
			close(stop)
			log.Println("Errors too match")
		}
	}
}

var stop = make(chan struct{})
var counter = &ErrorCounter{}
var wg sync.WaitGroup

// Hendler - functions hendler
func Hendler(tasks []function, n int) {
	for _, task := range tasks {
		wg.Add(1)
		go func(task function) {
			maker(task,n)
			wg.Done()
		}(task)
	}
	wg.Wait()
}
