package auth

import (
	"net/http"
	"os"
)

const htmlPagePath = "./mc/auth/authComplete.html"

func (as *AuthSession) AuthCompleteHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: check that this was a redirection from token handler.
	pageData, err := os.ReadFile(htmlPagePath)
	if err != nil {
		as.Logger.Debug("AuthCompleteHandler", "err", err)
		w.Write([]byte("Authentication complete. You may now close this page."))
	}
	w.Write(pageData)
}
