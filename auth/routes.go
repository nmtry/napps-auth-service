package auth

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/register", handleRegister).Methods("POST")

	return r
}
