package middleware

import "net/http"

func AuthStateHandler(state string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if state != r.URL.Query().Get("state") {
			http.Error(w, "Invalid state detected", http.StatusBadRequest)
			return
		}
		return
	}
}
