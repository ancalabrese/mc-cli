package auth

import (
	"context"
	"fmt"
	"net/http"
	url "net/url"
	"time"

	"github.com/ancalabrese/mc-cli/mc/auth/middleware"
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/ancalabrese/mc-cli/utils"
	"golang.org/x/oauth2"
)

const (
	mcUrlAuthorizationPath = "oauth/authorize"
	mcTokenUrlPath         = "token"
	mcApiUrlPath           = "api"
)

type AuthorizationResponseType int

var AuthResponseTypeCode oauth2.AuthCodeOption = oauth2.SetAuthURLParam("response_type", "code")
var AuthResponseTypeToken oauth2.AuthCodeOption = oauth2.SetAuthURLParam("response_type", "token")

type authSession struct {
	authState                 string
	authorizationCompleteChan <-chan struct{}
	authSever                 *http.Server
}

func Login(ctx context.Context, c *config.Config) error {
	authContext, cancel := context.WithCancel(ctx)
	defer cancel()

	authSession := &authSession{}
	authSession.initAuthServer()

	go authSession.startAuthenticationServer(authContext)
	if err := RequestAuthCode(c); err != nil {
		return err
	}

	<-authSession.authorizationCompleteChan
	return nil
}

func RequestAuthCode(c *config.Config) error {
	addr, err := c.Authentication.Get(config.McHostKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve auth address: %w", err)
	}

	clientId, err := c.Authentication.Get(config.McClientIdKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve client id: %w", err)
	}

	clientSecret, err := c.Authentication.Get(config.McSecretKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve client secret : %w", err)
	}

	callbackUrl, err := c.Authentication.Get(config.McCallbackUriKey)
	if err != nil {
		return fmt.Errorf("Failed to retrieve auth callback url: %w", err)
	}

	mcAuthUrl, err := url.JoinPath(addr, mcUrlAuthorizationPath)
	mcTokenUrl, err := url.JoinPath(addr, mcApiUrlPath, mcTokenUrlPath)
	utils.Check(err)

	oauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:   mcAuthUrl,
			TokenURL:  mcTokenUrl,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	authCodeUrl := oauthConfig.AuthCodeURL(
		oauth2.GenerateVerifier(),
		oauth2.AccessTypeOffline,
		AuthResponseTypeCode)

	// redirect_url breaks Mobicontrol authorization page
	parsedUrl, err := url.Parse(authCodeUrl)
	utils.Check(err)
	urlParams := parsedUrl.Query()
	urlParams.Del("redirect_uri")
	parsedUrl.RawQuery = urlParams.Encode()

	fmt.Printf("Login at:\n%s\n", parsedUrl.String())
	return nil
}

func (as *authSession) initAuthServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/callback", middleware.AuthStateHandler(as.authState))

	as.authSever = &http.Server{
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (as *authSession) startAuthenticationServer(ctx context.Context) {
	go func() {
		utils.Check(ctx.Err())
		err := as.authSever.ListenAndServe()
		utils.Check(err)
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := as.authSever.Shutdown(shutdownCtx)
	utils.Check(err)
}
