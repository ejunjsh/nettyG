package gonet


type Event int

const(
	CONNECTED Event=iota
	READABLE
	ERROR
)


type Events struct {
	fs map[Event]func(BufferChannel)
	fss []Event
}

func Register(e Event, f func(BufferChannel)){
    fs[e]=f
}

func trigger(e Event){
	switch e {
	case CONNECTED:
	case READABLE:
	case ERROR:
	default:
		return
	}
}