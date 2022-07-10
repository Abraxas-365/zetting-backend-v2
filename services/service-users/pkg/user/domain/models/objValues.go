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
