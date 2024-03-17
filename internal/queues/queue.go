package queues

type Queue interface {
	Init()
	AddConsumer(f func(e Event))
	Publish(message []byte)
	Close()
}

type Event interface {
	GetEventUUID() string
}
