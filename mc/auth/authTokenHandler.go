package auth

import (
	"context"
	"net/http"
)

func (as *authSession) OAuthTokenExchangeHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t, err := as.oauthConfig.Exchange(ctx, as.authorizationCode)
			if err != nil {
				as.Logger.Error("Token exchange failed", "err", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			as.Token = t
			as.tokenStore.Save(as.oauthConfig.ClientID, t)
			as.Logger.Debug("Access token received")

			http.Redirect(w, r, "/complete", http.StatusFound)
		})
}
