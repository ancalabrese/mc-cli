package auth

import (
	"net/http"
)

func (as *authSession) AuthStateHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			got := r.URL.Query().Get("state")
			want := as.authState
			if as.authState == "" || want != got {
				as.Logger.Error("Auth state mismatch", "IP", r.RemoteAddr)
				as.Logger.Debug("State", "want", as.authState, "got", got)
				http.Error(w, "Invalid auth state", http.StatusBadRequest)
				return
			}

			as.Logger.Debug("Authentication state: OK")
			// State is unique for each auth request
			as.authState = ""
			next.ServeHTTP(w, r)
		})
}
