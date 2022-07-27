package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"service-users/pkg/user/domain/events"
)

func (mq *mqPublisher) PublishEvents(events events.EventsInfo) error {
	for _, e := range events {
		fmt.Println(e.Event.Name())
		eventBytes := new(bytes.Buffer)
		json.NewEncoder(eventBytes).Encode(e)
		if err := mq.publisher.PublishToExchange(e.Event.Exchange(), e.Event.Routing(), eventBytes.Bytes()); err != nil {
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
