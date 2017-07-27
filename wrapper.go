package gonet

import "bytes"

type BufferChannel struct{
	bytes.Buffer
}


func (b *BufferChannel) Read(){

}