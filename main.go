package main

import (
	"napps-auth-service/auth"
	"net/http"
)

func main() {
	r := auth.InitRoutes()

	http.ListenAndServe(":8080", r)
}
