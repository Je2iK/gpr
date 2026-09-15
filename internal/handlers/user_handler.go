package handlers

import (
	"encoding/json"
	"net/http"
	"practice/internal/models"
	"practice/internal/service"
)

type UserHandler struct {
	s *service.userService
}
func NewUserHandler(s *service.userService) *UserHandler{
	return &UserHandler{s : s}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	var req models.CreateUserParams
	if err:=json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err:= h.s.Register(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id" : userId})
}