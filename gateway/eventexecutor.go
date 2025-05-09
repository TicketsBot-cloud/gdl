package gateway

import (
	"encoding/json"
	"reflect"

	"github.com/TicketsBot-cloud/gdl/gateway/payloads/events"
	"github.com/sirupsen/logrus"
)

func (s *Shard) ExecuteEvent(eventType events.EventType, data json.RawMessage) {
	dataType := events.EventTypes[eventType]
	if dataType == nil {
		return
	}

	event := reflect.New(dataType)
	if err := json.Unmarshal(data, event.Interface()); err != nil {
		logrus.Warnf("error whilst decoding event data: %s", err.Error())
	}

	for _, listener := range s.ShardManager.EventBus.Listeners {
		fn := reflect.TypeOf(listener)
		if fn.NumIn() != 2 {
			continue
		}

		ptr := fn.In(1)
		if ptr.Kind() != reflect.Ptr {
			continue
		}

		if ptr.Elem() == dataType {
			reflect.ValueOf(listener).Call([]reflect.Value{
				reflect.ValueOf(s),
				event,
			})
		}
	}
}
