package tasks

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestHendler(t *testing.T) {
	tasks := []function{
		func() error {
			time.Sleep(time.Second)
			return errors.New("some error")
		
		},
		func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("some job 2")
			return nil
		},
		func() error {
			time.Sleep(3 * time.Second)
			fmt.Println("some job 3")
			return nil
		},

		func() error {
			time.Sleep(5 * time.Second)
			fmt.Println("some job 4")
			return nil
		},
		func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("some job 5")
			return nil
		
		},
	}
	Hendler(tasks, 1)
}
