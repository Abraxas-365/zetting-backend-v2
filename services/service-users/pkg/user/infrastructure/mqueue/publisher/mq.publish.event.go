package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"service-users/pkg/user/domain/events"
)

func (mq *mqPublisher) PublishEvent(event events.EventInfo) error {

	eventBytes := new(bytes.Buffer)
	json.NewEncoder(eventBytes).Encode(event)
	if err := mq.publisher.PublishToExchange(event.Event.Exchange(), event.Event.Routing(), eventBytes.Bytes()); err != nil {
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
		return err
	}

	return nil
}
