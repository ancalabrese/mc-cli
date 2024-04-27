package auth

import (
	"net/http"
	"os"
)

const htmlPagePath = "./mc/auth/authComplete.html"

func (as *authSession) AuthCompleteHandler(w http.ResponseWriter, r *http.Request) {
	if as.Token == nil {
		as.Logger.Error("AuthCompleteHandler: user got to complete without a valid access token")
		http.Error(w, "Invalid access token", http.StatusBadRequest)
		return
	}
	pageData, err := os.ReadFile(htmlPagePath)
	if err != nil {
		as.Logger.Debug("AuthCompleteHandler", "err", err)
		w.Write([]byte("Authentication complete. You may now close this page."))
	}
	w.Write(pageData)
}
