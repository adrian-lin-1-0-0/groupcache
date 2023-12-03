package discovery

type Client interface {
	GetServices(string) ([]Service, error)
	WatchPrefix(string) (EventChan, error)
	Register(Service) error
	Unregister(Service) error
}

type EventChan <-chan Event

type Service struct {
	Name     string
	Endpoint string
}

type Event struct {
	Type  EventType
	Key   string
	Value string
}

type EventType int

const (
	EventPut EventType = iota
	EventDelete
)
