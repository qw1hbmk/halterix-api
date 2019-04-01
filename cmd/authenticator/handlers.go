package authenticator

import (
	"net/http"
	"strings"
)

// AuthMiddleware is the auth middleware layer.
// It will check for an authorization header
func APIKeyMiddleware(h http.Handler, db *database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		// Only expecting two fields in the header
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "API-Key" {
			http.Error(w, "No Authorization", http.StatusUnauthorized)
			return

		} else if Validate(auth[1], db) == false {
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
