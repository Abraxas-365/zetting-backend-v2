package models

type Creator struct {
	IsCreator  bool     `bson:"is_creator" json:"is_creator"`
	Skills     []string `bson:"skills" json:"skills,omitempty"`
	Categories []string `bson:"categories" json:"categories,omitempty"`
}
