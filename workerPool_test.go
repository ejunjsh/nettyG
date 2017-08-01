package gonet

import (
	"testing"
	"fmt"
	"time"
)

func TestWorkerPool_Execute(t *testing.T) {
	wp:=NewWorkerPool(10)
	wp.Execute(func() {
		fmt.Println(1)
	})

	wp.Start()

	wp.Execute(func() {
		fmt.Println(2)
	})

	wp.Execute(func() {
		fmt.Println(3)
	})

	time.Sleep(5*time.Second)

	wp.Stop()

	time.Sleep(20*time.Second)

}
