package auth

import (
	"context"
	"fmt"
	"net/http"
	url "net/url"
	"time"

	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
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

type AuthSession struct {
	authState                 string
	authorizationCode         string
	oauthConfig               *oauth2.Config
	tokenStore                TokenStore
	authorizationCompleteChan chan *oauth2.Token
	authSever                 *http.Server
	Logger                    hclog.Logger
}

func NewAuthSession(l hclog.Logger) *AuthSession {
	return &AuthSession{
		Logger:     l,
		tokenStore: &KeyringTokenStore{},
	}

}

func (as *AuthSession) Login(ctx context.Context, c *config.Config) error {
	authContext, cancel := context.WithCancel(ctx)
	defer cancel()

	as.initAuthServer(authContext)

	go as.startAuthenticationServer(authContext)
	as.authState = oauth2.GenerateVerifier()

	if err := as.requestAuthCode(c); err != nil {
		return err
	}

	<-as.authorizationCompleteChan
	return nil
}

func (as *AuthSession) requestAuthCode(c *config.Config) error {
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

	as.oauthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:   mcAuthUrl,
			TokenURL:  mcTokenUrl,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	authCodeUrl := as.oauthConfig.AuthCodeURL(
		as.authState,
		oauth2.AccessTypeOffline)

	// redirect_url breaks Mobicontrol authorization page
	parsedUrl, err := url.Parse(authCodeUrl)
	utils.Check(err)
	urlParams := parsedUrl.Query()
	urlParams.Del("redirect_uri")
	parsedUrl.RawQuery = urlParams.Encode()

	fmt.Printf("Login at:\n%s\n", parsedUrl.String())
	return nil
}

func (as *AuthSession) initAuthServer(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/host", as.AuthStateHandler(as.OAuthTokenExchangeHandler(ctx, as.authorizationCode)))

	as.authSever = &http.Server{
		Addr:           "localhost:8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       as.Logger.StandardLogger(&hclog.StandardLoggerOptions{}),
	}
}

func (as *AuthSession) startAuthenticationServer(ctx context.Context) {
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
