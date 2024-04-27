package auth

import "net/http"

func (as *AuthSession) AuthorizationCodeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.URL.Query().Get("code")
		if c == "" {
			as.Logger.Error("AuthorizationCodeHandler: Empty authorization code.")
			http.Error(w, "Invalid authorization code", http.StatusBadRequest)
			return
		}
		as.authorizationCode = c
		as.Logger.Debug("AuthorizationCodeHandler: Authorization code ok")
		next.ServeHTTP(w, r)
	})
}
