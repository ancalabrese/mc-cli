package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

const successHtmlPagePath = "./authComplete.html"

func (as *authSession) OAuthTokenExchangeHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t, err := as.oauthConfig.Exchange(ctx, as.authorizationCode)
			if err != nil {
				as.Logger.Error("Token exchange failed", "err", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			as.Logger.Debug("Access token received")
			as.apiSecretStore.SaveRefreshAccessToken(as.oauthConfig.ClientID, t.RefreshToken)

			pageData, err := os.ReadFile(successHtmlPagePath)
			if err != nil {
				as.Logger.Debug("AuthCompleteHandler", "err", err)
				w.Write([]byte("Authentication complete. You may now close this page."))
			}
			w.Write(pageData)

			as.Logger.Info("Authentication complete! You can run other commands now")
			fmt.Println("Authentication complete! You can run other commands now")
			as.authorizationCompleteChan <- t
		})
}
