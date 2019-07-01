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

// func cancelled() bool {
// 	select {
// 	case <-stop:
// 		return true
// 	default:
// 		return false
// 	}
// }

func maker(task function, countErrors int) {
    if err := task(); err != nil {
		count := counter.Add()
		if count >= countErrors {
			close(stop)
			log.Println("Errors too match")
		}
	}
}

var stop = make(chan struct{})
var counter = &ErrorCounter{}
var wg sync.WaitGroup

// Hendler - functions hendler
func Hendler(tasks []function, countWokers int, countErrors int) {
	wokers := make(chan struct{},countWokers)
	done := make(chan struct{},countWokers)
	for _, task := range tasks {
		wokers <- struct{}{}
		go func(task function) {
			task()
			<-wokers
			done <- struct{}{}
			//maker(task,countErrors)
		}(task)
	}

	log.Println("!!!!!!!!!!!!!!!!!!!!!!!!")

	close(wokers)	
	
	for res := range done {
		log.Println(res)
	}
}
