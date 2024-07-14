package controller

import (
	"awesomeProject/internal/entity"
	"awesomeProject/internal/service"
	"encoding/json"
	"net/http"
)

type Controller interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type con struct {
	service.Service
}

func NewController(service2 service.Service) Controller {
	return &con{service2}
}

func (c *con) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("Id") //need read query
	user, err := c.Service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
func (c *con) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	updatedUser, err := c.Service.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
func (c *con) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = c.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
func (c *con) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("Id") //need read query
	err := c.Service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
