package netgo



func handle(chl *channel){
	chl.runReadEventLoop()
	chl.runWriteEventLoop()
}