package service

import (
	"fmt"
	"service-users/pkg/user/domain/models"
)

func (s *userService) CanUserBeCreated(new models.User) (bool, error) {
	//TODO: applay the logic of chaeck email , check pass , etc
	if !new.IsValid() {
		fmt.Println("ES INVALIDO EL PASS")
		return false, nil
	}
	_, exist, _ := s.userRepo.GetPrivateUser("contact.email", new.Contact.Email)
	if exist {
		fmt.Println("EXISTEEE")
		return false, nil
	}

	fmt.Println("SI SE PUEDE ")
	return true, nil
}
