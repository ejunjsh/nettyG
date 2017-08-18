package netgo

type Encoder interface {
	onWrite(data interface{}) interface{}
}


type Decoder interface {
	onRead(data interface{}) interface{}
}


type LineCodec struct {

}


func (l *LineCodec) onWrite(data interface{}) interface{} {
	return nil
}

func (l *LineCodec) onRead(data interface{}) interface{}{
 	return  nil
}