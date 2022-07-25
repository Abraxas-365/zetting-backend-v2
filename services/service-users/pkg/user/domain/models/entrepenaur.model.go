package models

type Entrepenaur struct {
	IsEntrepenaur   bool   `bson:"is_entrepenaur" json:"is_entrepenaur"`
	Business        string `bson:"business" json:"business,omitempty"`
	Ruc             string `bson:"ruc" json:"ruc,omitempty"`
	BusinessAddress string `bson:"business_address" json:"business_address,omitempty"`
}
