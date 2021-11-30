package dispatcher

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/kiem-toan/pkg/errorx"

	"github.com/kiem-toan/pkg/event/listener"
)

// Như là 1 nhà điều phối event
// Chỉ có 1 dispatcher trong chương trình
type Dispatcher struct {
	// 1 event sẽ đi theo nhiều listener (process manager)
	events map[reflect.Type][]listener.Listener
}

// Khởi tạo nhà điều phối Dispatcher
func NewDispatcher() *Dispatcher {
	d := &Dispatcher{}
	d.events = make(map[reflect.Type][]listener.Listener)
	return d
}

func (d *Dispatcher) AddEventListner(eventListener listener.Listener, event interface{}) error {
	eventType := reflect.TypeOf(event)
	_, exists := d.events[eventType]
	if !exists {
		d.events[eventType] = make([]listener.Listener, 0)
	}
	d.events[eventType] = append(d.events[eventType], eventListener)
	return nil
}

func (d *Dispatcher) Dispatch(event interface{}) error {
	eventType := reflect.TypeOf(event)
	if _, ok := d.events[eventType]; !ok {
		return errorx.Errorf(http.StatusInternalServerError, nil, fmt.Sprintf("The '%s' event is not registered", eventType.String()))
	}
	listens := d.events[eventType]
	for _, listen := range listens {
		listen.Listen(event)
	}
	return nil
}
