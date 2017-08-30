package nettyG

import (
	"bytes"
	"strings"
)

type LineCodec struct {
	delimiter string
	Codec
}

func NewLineCodec(delimiter string) *LineCodec {
	return &LineCodec{delimiter:delimiter}
}

func (l *LineCodec) ChannelRead(c *HandlerContext,data interface{}) error{
	if buffer,ok:=data.(*bytes.Buffer);ok{
		for {
			s := buffer.String()
			r := []rune(s)
			index := strings.Index(s, l.delimiter)
			if index > 0 {
				r = r[0:index]
				b := []byte(string(r))
				buffer.Next(len(b))
				buffer.Next(len(l.delimiter))
				c.FireChannelRead(b)
			} else {
				break
			}
		}

	}
	return nil
}

func (l *LineCodec) Write(c *HandlerContext,data interface{}) error{
	if b,ok:=data.([]byte);ok{
		c.Write(append(b, []byte(l.delimiter)...))
	}
	return nil
}

