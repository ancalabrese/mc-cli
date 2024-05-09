package auth

import (
	"fmt"
	"net/http"
	"os"
)

const htmlPagePath = "./mc/auth/authComplete.html"

func (as *authSession) AuthCompleteHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if as.Token == nil {
				as.Logger.Error("AuthCompleteHandler: no valid access token")
				http.Error(w, "Invalid access token", http.StatusBadRequest)
				return
			}
			pageData, err := os.ReadFile(htmlPagePath)
			if err != nil {
				as.Logger.Debug("AuthCompleteHandler", "err", err)
				w.Write([]byte("Authentication complete. You may now close this page."))
			}
			w.Write(pageData)
			as.Logger.Info("Authentication complete! You can run other commands now")
			fmt.Println("Authentication complete! You can run other commands now")
			as.authorizationCompleteChan <- as.Token
		})
}
