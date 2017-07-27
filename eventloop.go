package gonet

import "sync"

type EventLoop struct {
    sync.Mutex
	events []interface{}
}


func newEventLoop() *EventLoop{
	return &EventLoop{events:[]interface{}{}}
}

func (el *EventLoop) put(event interface{}){
	el.Lock()
	defer el.Unlock()
	el.events=append(el.events,event)
}


func (el *EventLoop) run(){
	for{
		el.Lock()
		e:=el.events[0]
		el.events=el.events[1:]
		if f,ok:= e.(func (*Context));ok{
			f()
		}
		el.Unlock()

	}
}