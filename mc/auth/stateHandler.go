package auth

import (
	"net/http"
)

func (as *AuthSession) AuthStateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		got := r.URL.Query().Get("state")
		want := as.authState
		if as.authState == "" || want != got {
			as.Logger.Error("Auth state mismatch", "IP", r.RemoteAddr)
			as.Logger.Debug("State", "want", as.authState, "got", got)
			http.Error(w, "Invalid auth state", http.StatusBadRequest)
			return
		}

		as.Logger.Debug("Valid auth state")
		// State is unique for each auth request
		as.authState = ""
		return
	}
}
