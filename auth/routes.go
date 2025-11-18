package auth

import (
	"napps-auth-service/auth/error"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(LoggingMiddleware, error.RecoveryMiddleware)

	/* YOU CAN ALSO SET MIDDLEWARE PER ROUTE
		// Public route
	    r.HandleFunc("/public", HelloHandler).Methods("GET")

	    // Protected route with middleware
	    r.Handle("/secure", AuthMiddleware(http.HandlerFunc(HelloHandler))).Methods("GET")
	*/

	r.HandleFunc("/auth/register", handleRegister).Methods("POST")
	r.HandleFunc("/auth/authorize", handleRegister).Methods("POST")
	r.HandleFunc("/auth/tokens", handleRegister).Methods("POST")
	r.HandleFunc("/auth/refresh", handleRegister).Methods("POST")

	return r
}
