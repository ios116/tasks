package tasks

import (
	"log"
	"sync"
)

// function - it's task signature
type function func() error

// Counter - error, job counter
type Counter struct {
	sync.Mutex
	errorCount int
	jobCount   int
}

// addError - safely adds one to errors counter
func (ec *Counter) addError() {
	ec.Lock()
	ec.errorCount++
	ec.Unlock()
}

// getError - errors counter is safe for concurrent use
func (ec *Counter) getError() int {
	ec.Lock()
	defer ec.Unlock()
	return ec.errorCount
}

// addJob - safely adds one to job counter
func (ec *Counter) addJob() {
	ec.Lock()
	ec.jobCount++
	ec.Unlock()
}

// getJob - job counter is safe for concurrent use
func (ec *Counter) getJob() int {
	ec.Lock()
	defer ec.Unlock()
	return ec.jobCount

}

// cancelled - cancel if channel "done" is closed
func cancelled(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// maker - function executer of the task if erros more then countErrors then done channel is closed
func maker(task function, countErrors int, done chan struct{}) {
	if err := task(); err != nil {
		counter.addError()
		if counter.getError() >= countErrors {
			close(done)
			log.Println("Errors too match")
		}
	} else {
		counter.addJob()
	}
}

var counter = &Counter{}

// Hendler -  Task handler returns error counter and completed task counter
func Hendler(tasks []function, countWokers int, countErrors int) (int, int) {
	wokers := make(chan struct{}, countWokers)
	done := make(chan struct{}, countWokers)
	var wg sync.WaitGroup
	for _, task := range tasks {
		wokers <- struct{}{}
		wg.Add(1)
		if cancelled(done) {
			return counter.getError(), counter.getJob()
		}
		go func(task function) {
			maker(task, countErrors, done)
			<-wokers
			wg.Done()
		}(task)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	for res := range done {
		log.Println(res)
	}
	return counter.getError(), counter.getJob()
}
