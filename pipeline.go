package gonet


type Pipeline struct {
	handlers []Handler
}

func NewPipeline() *Pipeline{
	return &Pipeline{[]Handler{}}
}

func (p *Pipeline)Register(handler Handler)  {
	p.handlers=append(p.handlers,handler)
}

