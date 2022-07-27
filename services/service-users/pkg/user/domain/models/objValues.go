package models

//Obj Password

type Password string

func (p Password) Validate() bool {
	//TODO: make more validations
	switch true {
	case len(p) < 7:
		return false
	}
	return true
}

func (p Password) EqualTo(other Password) bool {
	return p == other
}

//Obj Contact
//
type Contact struct {
	Phone string `bson:"phone" json:"phone"`
	Email string `bson:"email" json:"email"`
}

type Name struct {
	FirstName  string `bson:"first_name" json:"first_name,omitempty"`
	MiddleName string `bson:"middle_name" json:"middle_name,omitempty"`
	LastName   string `bson:"last_name" json:"last_name"`
}
