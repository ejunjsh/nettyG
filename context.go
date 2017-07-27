package gonet

import "bytes"

type Context struct {
	buffer *bytes.Buffer
}

func newContext() *Context{
	return &Context{ bytes.NewBuffer( make([]byte,1024))}
}

func (c *Context) Write(data interface{}){

}

func (c *Context) Read(data interface{}){

}
