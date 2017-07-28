package gonet


type Context struct {
    data interface{}
}

func newContext() *Context{
	return &Context{}
}

func (c *Context) Write(data interface{}){
    c.data=data
}

