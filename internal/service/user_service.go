package service

import (
	"context"
	"practice/internal/models"
	"practice/internal/repository"
	"log"
)

type UserService interface {
	Register(ctx context.Context, params models.CreateUserParams) (int, error)
	Logout(ctx context.Context, params models.CreateUserParams) (int, error)
}
type userService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r: r}
}
func (s *userService) Register(ctx context.Context, params models.CreateUserParams) (int, error) {
	name, err := s.r.GetUserIdByName(ctx, params.Name)
	if err != nil{
		log.Println("User already is: %s, %v", name, err)
		return 0, nil
	}
	return s.r.CreateUser(ctx, params.Name, params.Password)
}
func (s *userService) Logout(ctx context.Context, params models.CreateUserParams) (int, error){
	return s.r.DeleteUserByName(ctx, params.Name)
}