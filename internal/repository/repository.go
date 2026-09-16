package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)


type UserRepository interface {
	CreateUser(ctx context.Context, name string, password string) (int, error)
	GetUserById(ctx context.Context, id int) (string, error)
	GetUserIdByName (ctx context.Context, name string) (int, error)
	DeleteUserByName(ctx context.Context, name string) (int, error)
}
type userRepository struct{
	pool *pgxpool.Pool
}
func NewUserRepository(pool *pgxpool.Pool) UserRepository{
	return &userRepository{
		pool: pool,
	}
}
func (r *userRepository) CreateUser(ctx context.Context, name string, password string) (int, error){
	query := `INSERT INTO users (name, password) VALUES($1, $2) RETURNING id;`
	var createdId int
	err := r.pool.QueryRow(ctx, query, name, password).Scan(&createdId)
	if err != nil{
		return 0, err

	}
	return createdId, nil

}
func (r *userRepository) GetUserById(ctx context.Context, id int) (string, error){
	query := `SELECT name FROM users WHERE id = $1;`

	var name string
	err := r.pool.QueryRow(ctx, query, id).Scan(&name)
	if err != nil{
		return "", fmt.Errorf("Do sviyazi: %w", err)
	}
	return name, nil

}
func (r *userRepository) GetUserIdByName(ctx context.Context, name string) (int, error){
	query := `SELECT id FROM users WHERE name = $1;`
	var id int
	err := r.pool.QueryRow(ctx, query, name).Scan(&id)
	if err != nil{
		return 0, fmt.Errorf("Do svyz: %w", err)
	}
	return id, nil
}
func (r *userRepository) DeleteUserByName(ctx context.Context,name string) (int, error){
	query := `DELETE FROM users WHERE name = $1 RETURNING id;`
	var userId int
	err := r.pool.QueryRow(ctx, query, name).Scan(&userId)
	if err != nil{
		return 0, fmt.Errorf("Not find user: %w", err)
	}
	return userId, nil
}