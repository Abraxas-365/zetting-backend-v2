package models

type Entrepenaur struct {
	IsEntrepenaur   bool   `bson:"is_entrepenaur" json:"is_entrepenaur"`
	Business        string `bson:"business,omitempty" json:"business,omitempty"`
	Ruc             string `bson:"ruc,omitempty" json:"ruc,omitempty"`
	BusinessAddress string `bson:"business_address,omitempty" json:"business_address,omitempty"`
}
