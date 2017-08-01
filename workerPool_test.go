package gonet

import (
	"testing"
	"sync"
	"fmt"
)

func workerpool() {
	wg := new(sync.WaitGroup)
	wp:=NewWorkerPool(4)
	wp.Start()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		wp.Execute(func() {
			for j := 0; j < 20000; j++ {

			}
			wg.Done()
		})
	}

	wg.Wait()
}

func nopool() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			for j := 0; j < 20000; j++ {

			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func gopool() {
	wg := new(sync.WaitGroup)
	data := make(chan int, 100)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for _ = range data {
				for j := 0; j < 20000; j++ {

				}
			}
		}(i)
	}

	for i := 0; i < 10000; i++ {
		data <- i
	}
	close(data)
	wg.Wait()
}

func BenchmarkWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		workerpool()
	}
}

func BenchmarkNopool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nopool()
	}
}

func BenchmarkGopool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopool()
	}
}

func TestNewWorkerPool(t *testing.T) {
	wg := new(sync.WaitGroup)
	wp:=NewWorkerPool(4)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		wp.Execute(func() {
			fmt.Println(1)
			wg.Done()
		})

	}
	wp.Start()
	wg.Wait()
	wp.Stop()
}