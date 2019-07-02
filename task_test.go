package tasks

import (
	"errors"
	"testing"
	"time"
)

func TestHendler(t *testing.T) {

	functions := make([]function, 10)
	for i := range functions {
		n := i
		functions[i] = func() error {
			t.Logf("start job %d", n)
			time.Sleep(3 * time.Second)
			t.Logf("== end == job %d", n)
			return nil
		}
	}

	functionsWitError := make([]function, 2)
	for i := range functionsWitError {
		n := i
		functionsWitError[n] = func() error {
			t.Logf("start job with err %d", n)
			time.Sleep(4 * time.Second)
			t.Logf("== end with error== job %d", n)
			return errors.New("functions with error")
		}
	}
	tasks := []function{
		functions[0],
		functions[1],
		functionsWitError[0],
		functions[3],
		functions[4],
		functions[5],
		functionsWitError[1],
		functions[7],
		functions[8],
		functions[9],
	}

	errCount, jobCount := Hendler(tasks, 3, 1)
	if errCount != 1 {
		t.Fatalf("we have ERRORS %d must have 1", errCount)
	}
	if jobCount != 2 {
		t.Fatalf("we have complited JOB %d must have 2", jobCount)
	}
}
