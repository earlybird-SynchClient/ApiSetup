package api

import (
	"net/http"
	"github.com/earlybird-SynchClient/ApiSetup/paola/auth"
)

func ForceLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := auth.Auth.GetCurrentUser(r)
		if currentUser == nil {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}