package auth

import (
	"context"
	"net/http"
)

func (as *AuthSession) OAuthTokenExchangeHandler(ctx context.Context, authCode string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := as.oauthConfig.Exchange(ctx, authCode)
		if err != nil {
			as.Logger.Error("Token exchange failed", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//TODO: save token from line 10

		http.Redirect(w, r, "/complete", http.StatusTemporaryRedirect)
	}
}
