package error

import (
	"fmt"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				details := fmt.Sprint(err)

				switch details {
				case "Email already registered":
					WriteError(w, http.StatusBadRequest, "Bad Request", details)
				default:
					WriteError(w, http.StatusInternalServerError, "Internal Server Error", details)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
