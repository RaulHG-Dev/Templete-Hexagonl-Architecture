package user

import (
	"architecture-hexagonal/internal/core/domain"
	"architecture-hexagonal/internal/ports/entities"
	"fmt"
	"net/http"
)

type Handler struct {
	UserService entities.UserService
}

func (h Handler) NewUserHandler(userService entities.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (usr Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user handler")
	user := domain.User{
		ID:    1,
		Name:  "test",
		Email: "mail@gmail.com",
	}
	fmt.Println("user handler", usr)

	id, err := usr.UserService.Save(user)
	fmt.Println("user handler", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created with ID: %d", id)))
	// w.Write([]byte(fmt.Sprintf("User created with ID: %d", id)))
}
