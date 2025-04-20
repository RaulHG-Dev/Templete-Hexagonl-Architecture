package user

import (
	"architecture-hexagonal/internal/core/domain"
	"fmt"
)

func (s Service) Create(user domain.User) (usr domain.User, err error) {
	// Here, we are calling the `Insert` method of the `UserRepository` interface to insert a user.
	id, err := s.Repo.Insert(user)
	fmt.Println(id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{}, nil
}
