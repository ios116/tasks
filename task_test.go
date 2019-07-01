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
			fmt.Println("start job 0")
			time.Sleep(5*time.Second)
			fmt.Println("stop job 0")
			//return errors.New("start error")
			
			return nil
		
		},

		func() error {
			fmt.Println("start job 1")
			time.Sleep(5 * time.Second)
			fmt.Println("stop job 1")
			return nil
		},


		func() error {
			fmt.Println("start job 2")
			time.Sleep(5 * time.Second)
			fmt.Println("stop job 2")
			return errors.New("start error")
		},
		func() error {
			fmt.Println("start job 3")
			time.Sleep(5 * time.Second)
			fmt.Println("stop job 3")
			
			return nil
		},

		func() error {
			fmt.Println("start job 4")
			time.Sleep(5 * time.Second)
			fmt.Println("stop job 4")
			
			return nil
		},
		func() error {
			fmt.Println("start job 5")
			time.Sleep(5 * time.Second)
			fmt.Println("stop job 5")
			
			return nil
		
		},
		func() error {
			fmt.Println("start job 6")
			time.Sleep(5 * time.Second)
			fmt.Println("stop job 6")
			
			return nil
		
		},
	}
	Hendler(tasks, 3,3)
}
