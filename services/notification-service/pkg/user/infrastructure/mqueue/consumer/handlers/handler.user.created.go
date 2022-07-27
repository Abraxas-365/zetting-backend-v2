package mqHandler

import (
	"encoding/json"
	"fmt"
	"notifications/pkg/user/domain/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) UserCreated(d amqp.Delivery) bool {

	body := struct {
		User models.User `json:"event"`
	}{}
	if err := json.Unmarshal([]byte(d.Body), &body); err != nil {
		return false
	}
	fmt.Println(body)

	if err := h.application.CreateUser(*&body.User); err != nil {
		return false
	}
	return true
}
