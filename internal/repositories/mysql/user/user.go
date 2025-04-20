package user

import (
	"architecture-hexagonal/internal/core/domain"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	// The `client` field is a pointer to a `gorm.DB` instance, which represents the database connection.
	Client *gorm.DB
}

func (r Repository) Insert(user domain.User) (id uint, err error) {
	fmt.Println("insert user")
	return 0, nil
}

func (r Repository) GetByID(id uint) (user domain.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Delete(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Update(user domain.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetAll() (users []domain.User, err error) {
	//TODO implement me
	panic("implement me")
}
