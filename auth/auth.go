package auth

import (
	"context"
	"fmt"
	"net/http"
	url "net/url"
	"time"

	"github.com/ancalabrese/mc-cli/config"
	"github.com/ancalabrese/mc-cli/storage"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
)

const (
	oauthPath         = "oauth"
	authorizationPath = "authorize"
	tokenUrlPath      = "token"
	apiPrefixPath     = "Mobicontrol"
	apiPrefixUrlPath  = "api"
)

var AuthResponseTypeCode oauth2.AuthCodeOption = oauth2.SetAuthURLParam("response_type", "code")
var AuthResponseTypeToken oauth2.AuthCodeOption = oauth2.SetAuthURLParam("response_type", "token")

type authSession struct {
	Logger                    hclog.Logger
	authState                 string
	authorizationCode         string
	oauthConfig               *oauth2.Config
	apiSecretStore            *storage.ApiSecretStore
	authorizationCompleteChan chan *oauth2.Token
	authServer                *http.Server
}

func NewAuthSession(ctx context.Context, c *config.Config, l hclog.Logger) (*oauth2.Token, error) {
	oauth2Config, err := GetOauth2Config(c)
	if err != nil {
		return nil, err
	}

	as := &authSession{
		Logger:                    l,
		apiSecretStore:            storage.NewApiSecretStore(l),
		authorizationCompleteChan: make(chan *oauth2.Token),
		oauthConfig:               oauth2Config,
	}
	token, err := as.login(ctx)
	return token, err
}

func GetOauth2Config(c *config.Config) (*oauth2.Config, error) {
	mcAuthUrl, err := url.JoinPath(c.Api.HostName, apiPrefixPath, oauthPath, authorizationPath)
	if err != nil {
		return nil, err
	}

	mcTokenUrl, err := url.JoinPath(c.Api.HostName, apiPrefixPath, apiPrefixUrlPath, tokenUrlPath)
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     c.Api.ClientId,
		ClientSecret: c.Api.ClientSecret,
		RedirectURL:  c.Api.CallbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:   mcAuthUrl,
			TokenURL:  mcTokenUrl,
			AuthStyle: oauth2.AuthStyleInHeader,
		},
	}, nil
}

func (as *authSession) login(ctx context.Context) (*oauth2.Token, error) {
	authContext, cancel := context.WithCancel(ctx)
	defer cancel()

	as.initAuthServer(authContext)

	go as.startAuthenticationServer(authContext)
	as.authState = oauth2.GenerateVerifier()

	if err := as.requestAuthCode(); err != nil {
		return nil, err
	}

	t := <-as.authorizationCompleteChan
	return t, nil
}

func (as *authSession) requestAuthCode() error {
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

func (as *authSession) initAuthServer(ctx context.Context) {
	mux := http.NewServeMux()
	mux.Handle("/callback", as.AuthStateHandler(
		as.AuthorizationCodeHandler(as.OAuthTokenExchangeHandler(ctx))))

	as.authServer = &http.Server{
		Addr:           "localhost:8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       as.Logger.StandardLogger(&hclog.StandardLoggerOptions{}),
	}
}

func (as *authSession) startAuthenticationServer(ctx context.Context) {
	go func() {
		utils.Check(ctx.Err())
		err := as.authServer.ListenAndServe()
		utils.Check(err)
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := as.authServer.Shutdown(shutdownCtx)
	utils.Check(err)
}
