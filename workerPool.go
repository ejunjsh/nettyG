package gonet

import (
	"sync"
	"fmt"
)

type task func()

type worker struct {
	stopC chan bool
}

type WorkerPool struct {
	num int
	sync.Mutex
	taskQ chan task
	workers []*worker
}

func NewWorkerPool(workerNum int,queueCap int) *WorkerPool{
	return &WorkerPool{num:workerNum,taskQ:make(chan task,queueCap),workers:make([]*worker,workerNum)}
}

func (wp *WorkerPool) Execute(t task){
	wp.taskQ<-t
}

func (wp *WorkerPool) Start() *WorkerPool{
	for i:=0;i<wp.num;i++{
		wp.workers[i]=&worker{ make(chan bool)}
		w:=wp.workers[i]
		go func(i int) {
			for {
				    stop:=false
					select {
					    case f:=<-wp.taskQ:
							f()
					    case stop=<-w.stopC:
						     break

					}

				if stop{
					break
				}
			}
			fmt.Println("stop")
		}(i)
	}
	return wp
}

//func (wp *WorkerPool) Sync(){
//	<-wp.stopC
//}

func (wp *WorkerPool) Stop(){
	for _,w:=range wp.workers{
		w.stopC<- true
	}
}