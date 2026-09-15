package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

)


func NewPool(ctx context.Context, cfg string) (*pgxpool.Pool, error){
	pool, err := pgxpool.New(ctx, cfg)
	if err != nil{
		return nil, fmt.Errorf("gg")
	}
	return pool, nil

}
