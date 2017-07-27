package gonet


type Pipeline struct {
	handlers []Handler
	encoders []Encoder
	decoders []Decoder
}

func NewPipeline() *Pipeline{
	return &Pipeline{[]Handler{},[]Encoder{},[]Decoder{}}
}

func (p *Pipeline)RegisterHandler(handler Handler)  {
	p.handlers=append(p.handlers,handler)
}


func (p *Pipeline)AddEncoder(decoder Decoder) {
	p.decoders=append(p.decoders,decoder)
}


func (p *Pipeline)AddDecoder(encoder Encoder) {
	p.encoders=append(p.encoders,encoder)
}
