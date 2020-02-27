package mysql

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	p := NewPool(3)

	task := NewTask(func() error {
		fmt.Print(time.Now())
		return nil
	})

	go func() {
		for {
			p.entryChannel <- task
		}
	}()

	p.runPool()
}
