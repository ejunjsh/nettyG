package gonet

import "sync"

type eventType int

const (
   channelActive eventType =iota
	errorCaught
	channelRead

)


type event struct {
	data interface{}
	et eventType
}

type EventLoop struct {
    sync.Mutex
	events []*event
	workerCount int
}


func newEventLoop(workerCount int) *EventLoop{
	return &EventLoop{workerCount:workerCount}
}

func (el *EventLoop) put(e *event){
	el.Lock()
	defer el.Unlock()
	el.events=append(el.events,e)
}


func (el *EventLoop) run(pl *Pipeline){
	for i:=0;i<el.workerCount;i++{
		go func() {
			for{
				el.Lock()
				e:=el.events[0]
				el.events=el.events[1:]
				for _,h:=range pl.handlers{
					switch e.et {
					case channelActive:
						h.channelActive()
					case channelRead:
						h.channelRead()
					case errorCaught:
						h.errorCaught()
					}
				}
				el.Unlock()

			}
		}()
	}

}