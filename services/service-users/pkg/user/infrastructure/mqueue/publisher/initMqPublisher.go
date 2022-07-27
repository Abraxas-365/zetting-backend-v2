package mqpublisher

import (
	"service-users/internal/rabbit"
	"service-users/pkg/user/domain/ports"
)

type mqPublisher struct {
	publisher rabbit.MQueue
}

func NewMQPublisher(mq rabbit.MQueue) ports.UserMQPublisher {
	return &mqPublisher{
		mq,
	}
}
