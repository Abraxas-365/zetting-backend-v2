package models

type Creator struct {
	IsCreator  bool     `bson:"is_creator" json:"is_creator"`
	Skills     []string `bson:"skills,omitempty" json:"skills,omitempty"`
	Categories []string `bson:"categories,omitempty" json:"categories,omitempty"`
}
