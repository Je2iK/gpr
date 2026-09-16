package models


type CreateUserParams struct {
	Name string `json:"name"`
	Password string `json:"password"`
}