package repositories

import "architecture-hexagonal/internal/core/domain"

type UserRepository interface {
	Insert(user domain.User) (id uint, err error)
	GetByID(id uint) (user domain.User, err error)
	Delete(id uint) (err error)
	Update(user domain.User) (err error)
	GetAll() (users []domain.User, err error)
}
