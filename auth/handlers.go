package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var registerRequest RegisterRequest
	json.NewDecoder(r.Body).Decode(&registerRequest)

	newUser := User{
		Id:          uuid.New().String(),
		Firstname:   registerRequest.Firstname,
		Lastname:    registerRequest.Lastname,
		Email:       registerRequest.Email,
		Username:    registerRequest.Username,
		Password:    HashPassword(registerRequest.Password),
		CreatedTime: time.Now(),
	}

	InsertUser(&newUser)

	retVal := RegisterResponse{
		Id: newUser.Id,
	}

	json.NewEncoder(w).Encode(retVal)
}
