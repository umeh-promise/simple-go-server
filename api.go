package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body to User struct
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	if err = insertUser(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func insertUser(newUser User) error {
	// input server
	if newUser.FirstName == "" {
		return errors.New("First name is required")
	}

	if newUser.LastName == "" {
		return errors.New("Last name is required")
	}

	// Storage validation
	for _, user := range users {
		if user.FirstName == newUser.FirstName && user.LastName == newUser.LastName {
			return errors.New("User already exists")
		}
	}

	users = append(users, newUser)
	return nil
}
