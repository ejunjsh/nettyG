package nettyG



func handle(chl *Channel){
	chl.runReadEventLoop()
	//chl.runWriteEventLoop()
}