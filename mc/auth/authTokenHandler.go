package auth

import (
	"context"
	"net/http"
)

func (as *AuthSession) OAuthTokenExchangeHandler(ctx context.Context, authCode string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := as.oauthConfig.Exchange(ctx, authCode)
		if err != nil {
			as.Logger.Error("Token exchange failed", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		go func() {
			// Error is ignored. Saved failed for some reason, but we can use the token
			_ = as.tokenStore.Save(as.oauthConfig.ClientID, t)
		}()

		as.authorizationCompleteChan <- t
		http.Redirect(w, r, "/complete", http.StatusTemporaryRedirect)
	}
}