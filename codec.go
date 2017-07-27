package gonet

type Encoder interface {
	onWrite(c *Context)
}


type Decoder interface {
	onRead(c *Context)
}


type LineCodec struct {

}


func (l *LineCodec) onWrite(c *Context){

}

func (l *LineCodec) onRead(c *Context){

}