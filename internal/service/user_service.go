package service

import (
	"context"
	"practice/internal/models"
	"practice/internal/repository"
)

type UserService interface{
	Register(ctx context.Context, params models.CreateUserParams) (int, error)	
}
type userService struct {
	r *repository.UserRepository
}
func NewUserService(r *repository.UserRepository) UserService{
	return &userService{r: r}
}
func (s *userService) Register(ctx context.Context, params models.CreateUserParams) (int, error){
	return s.r.CreateUser(ctx, params.Name, params.Password)
}	