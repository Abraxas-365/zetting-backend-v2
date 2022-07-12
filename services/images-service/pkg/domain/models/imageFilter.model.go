package models

type ImageFilter struct {
	Name string `bson:"name" json:"name"`
	Tag  string `bson:"tag" json:"tag"`
	//convert to uuid
	Owner  string `bson:"owner" json:"owner"`
	Format string `bson:"format" json:"format"`
}
