package nettyG



func handle(chl *channel){
	chl.runReadEventLoop()
	chl.runWriteEventLoop()
}