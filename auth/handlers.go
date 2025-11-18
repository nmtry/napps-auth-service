package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// ONLY DO OIDC FOR LOGIN, WHICH WOULD HIT THE AUTHORIZE. IF NOT REGISTERED, DIRECT TO A NAPPS REGISTER PAGE, AND THEY CAN REGISTER LIKE NORMAL
// for register OIDC
// query params: client_id, redirect_uri: some.com/callback, response_type=code, scope=openid profile email, code_verifier

// for login OIDC

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var registerRequest RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	callback := r.URL.Query().Get("callback")
	if callback != "" {
		http.Redirect(w, r, callback, http.StatusSeeOther)
	} else {
		// validate client_id, scope, state, response_type (code), code challenge
		// Sidenote: code challenge -- rp generates random long string in the frontend, computes SHA256 of that verifier, sends
		// it to this endpoint as code_challenge; this generates a code, stores it in a db with a short lived ttl with the code challenge, maybe user id?
	}

	// Check if email is already registered
	if IsEmailUsed(registerRequest.Email) {
		panic("Email already registered")
	}

	newUser := User{
		Id:          uuid.New().String(),
		Firstname:   registerRequest.Firstname,
		Lastname:    registerRequest.Lastname,
		Email:       registerRequest.Email,
		Username:    registerRequest.Username,
		Password:    HashPassword(registerRequest.Password),
		CreatedTime: time.Now(),
		Roles:       []string{"napps-user"},
	}

	InsertUser(&newUser)

	retVal := RegisterResponse{
		Id: newUser.Id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retVal)
}
