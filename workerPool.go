package gonet

import (
	"sync"
	"sync/atomic"
)

type task func()

type WorkerPool struct {
	num int
	sync.Mutex
	taskQ []task
	stopFlag int32
}

func NewWorkerPool(workerNum int) *WorkerPool{
	return &WorkerPool{num:workerNum,taskQ:[]task{}}
}

func (wp *WorkerPool) Execute(t task){
	wp.Lock()
	defer wp.Unlock()
	wp.taskQ=append(wp.taskQ,t)
}

func (wp *WorkerPool) Start() *WorkerPool{
	for i:=0;i<wp.num;i++{
		go func() {
			for   atomic.LoadInt32(&(wp.stopFlag)) == 0{
					if len(wp.taskQ)>0{
						wp.Lock()
						t:=wp.taskQ[0]
						wp.taskQ= wp.taskQ[1:]
						wp.Unlock()
						t()
					}
			}
		}()
	}
	return wp
}

//func (wp *WorkerPool) Sync(){
//	<-wp.stopC
//}

func (wp *WorkerPool) Stop(){
	atomic.StoreInt32(&wp.stopFlag,1)
}