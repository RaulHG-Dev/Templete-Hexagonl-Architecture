package entities

import "architecture-hexagonal/internal/core/domain"

type UserService interface {
	Save(user domain.User) (uint, error)
	FindByID(id uint) (domain.User, error)
	FindAll() ([]domain.User, error)
	Delete(id uint) error
	Update(user domain.User) error
}
