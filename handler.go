package gonet

type Handler interface {
	input(interface{}) interface{}
	output(interface{}) interface{}
}

