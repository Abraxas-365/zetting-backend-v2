package events

import "github.com/google/uuid"

type ImagePosted struct {
	UserId    uuid.UUID `json:"user_id"`
	ImageName string    `json:"image_name"`
	ImageId   uuid.UUID `json:"image_id"`
	ImageEtag string    `json:"image_etag"`
}

func (u ImagePosted) Name() string {
	return "event.image.posted"
}
func (u ImagePosted) Exchange() string {
	return "Image"
}
func (u ImagePosted) Routing() string {
	return "posted"
}
