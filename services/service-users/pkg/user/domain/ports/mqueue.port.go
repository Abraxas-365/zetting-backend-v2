package ports

import "service-users/pkg/user/domain/events"

type UserMQPublisher interface {
	PublishEvent(events.EventInfo) error
	PublishEvents(events.EventsInfo) error
}
