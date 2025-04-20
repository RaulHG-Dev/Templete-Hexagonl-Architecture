package user

import "architecture-hexagonal/internal/ports/repositories"

type Service struct {
	Repo repositories.UserRepository
}
