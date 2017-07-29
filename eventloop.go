package gonet

import "sync"

type event struct {
	data interface{}
	handler interface{}
}

type EventLoop struct {
    sync.Mutex
	events []*event
}


func newEventLoop() *EventLoop{
	return &EventLoop{}
}

func (el *EventLoop) put(e *event){
	el.Lock()
	defer el.Unlock()
	el.events=append(el.events,e)
}


func (el *EventLoop) run(){
	for{
		el.Lock()
		e:=el.events[0]
		el.events=el.events[1:]
		if f,ok:= e.handler.(func (*Context));ok{
			f()
		}
		el.Unlock()

	}
}