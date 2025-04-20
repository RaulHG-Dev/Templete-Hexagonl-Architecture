package user

import (
	"architecture-hexagonal/internal/core/domain"
	"architecture-hexagonal/internal/ports/repositories"
)

type Service struct {
	Repo repositories.UserRepository
}

func (s Service) Save(user domain.User) (uint, error) {
	id, err := s.Repo.Insert(user)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func (s Service) FindByID(id uint) (domain.User, error) {
	usr, err := s.Repo.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}

	return usr, nil
}

func (s Service) FindAll() ([]domain.User, error) {
	users, err := s.Repo.GetAll()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (s Service) Delete(id uint) error {
	err := s.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Update(user domain.User) error {
	err := s.Delete(user.ID)
	if err != nil {
		return err
	}
	return nil
}
