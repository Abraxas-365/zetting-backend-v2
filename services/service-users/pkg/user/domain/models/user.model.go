package models

import (
	"time"

	"github.com/google/uuid"
)

// UserPrivate
type User struct {
	ID          uuid.UUID   `bson:"_id,omitempty" json:"id"`
	Dni         string      `bson:"dni,omitempty" json:"dni,omitempty"`
	Name        string      `bson:"name" json:"name,omitempty"`
	Nickname    string      `bson:"nickname" json:"nickname"`
	Username    string      `bson:"username" json:"username"` //unicoooooo
	Password    Password    `bson:"password" json:"password"`
	Description string      `bson:"description" json:"description"`
	Headline    string      `bson:"headline" json:"headline"`
	PerfilImage string      `bson:"perfil_image" json:"perfil_image,omitempty"`
	Creator     Creator     `bson:"creator" json:"creator"`
	Entrepenaur Entrepenaur `bson:"entrepenaur" json:"entrepenaur"`
	Contact     Contact     `bson:"contact" json:"contact"`
	Gender      string      `bson:"gender" json:"gender,omitempty"`
	Age         int         `bson:"age" json:"age,omitempty"`
	Verified    bool        `bson:"verified" json:"verified,omitempty"`
	Created     time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated     time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type UserPublic struct {
	ID          uuid.UUID   `bson:"_id,omitempty" json:"id"`
	Name        string      `bson:"name" json:"name"`
	Nickname    string      `bson:"nickname" json:"nickname"`
	Username    string      `bson:"username" json:"username"`
	Description string      `bson:"description" json:"description"`
	Headline    string      `bson:"headline" json:"headline"`
	PerfilImage string      `bson:"perfil_image" json:"perfil_image,omitempty"`
	Creator     Creator     `bson:"creator" json:"creator"`
	Entrepenaur Entrepenaur `bson:"entrepenaur" json:"entrepenaur"`
	Gender      string      `bson:"gender" json:"gender,omitempty"`
	Age         int         `bson:"age" json:"age,omitempty"`
	Verified    bool        `bson:"verified" json:"verified,omitempty"`
}

func (u *User) New() *User {
	u.ID = uuid.New()
	return u
}

func (u *User) IsValid() bool {
	// switch true {
	// case len(u.Name) == 0 || u.Name == "":
	// 	return false
	// case len(u.Nickname) == 0:
	// 	return false
	// }

	if !u.Password.Validate() {
		return false
	}
	return true
}

func (u *User) ToPublic() UserPublic {
	return UserPublic{
		ID:          u.ID,
		Name:        u.Name,
		Nickname:    u.Nickname,
		Username:    u.Username,
		Description: u.Description,
		Headline:    u.Headline,
		PerfilImage: u.PerfilImage,
		Creator:     u.Creator,
		Entrepenaur: u.Entrepenaur,
		Gender:      u.Gender,
		Age:         u.Age,
		Verified:    u.Verified,
	}

}
