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
	oauthPath         = "oauth"
	authorizationPath = "authorize"
	tokenUrlPath      = "token"
	apiPrefixPath     = "Mobicontrol"
	apiPrefixUrlPath  = "api"
)

var AuthResponseTypeCode oauth2.AuthCodeOption = oauth2.SetAuthURLParam("response_type", "code")
var AuthResponseTypeToken oauth2.AuthCodeOption = oauth2.SetAuthURLParam("response_type", "token")

type authSession struct {
	authState                 string
	authorizationCode         string
	oauthConfig               *oauth2.Config
	tokenStore                TokenStore
	authorizationCompleteChan chan *oauth2.Token
	Token                     *oauth2.Token
	authSever                 *http.Server
	Logger                    hclog.Logger
}

func NewAuthSession(ctx context.Context, c *config.Config, l hclog.Logger) error {
	addr := c.Host.HostName
	clientId := c.Host.ClientId
	clientSecret := c.Host.ClientSecret
	callbackUrl := c.Host.CallbackURL

	mcAuthUrl, err := url.JoinPath(addr, apiPrefixPath, oauthPath, authorizationPath)
	mcTokenUrl, err := url.JoinPath(addr, apiPrefixPath, apiPrefixUrlPath, tokenUrlPath)
	if err != nil {
		return err
	}

	as := &authSession{
		Logger:                    l,
		tokenStore:                &KeyringTokenStore{},
		authorizationCompleteChan: make(chan *oauth2.Token),
		oauthConfig: &oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			RedirectURL:  callbackUrl,
			Endpoint: oauth2.Endpoint{
				AuthURL:   mcAuthUrl,
				TokenURL:  mcTokenUrl,
				AuthStyle: oauth2.AuthStyleInHeader,
			},
		},
	}
	return as.login(ctx)
}

func (as *authSession) login(ctx context.Context) error {
	authContext, cancel := context.WithCancel(ctx)
	defer cancel()

	as.initAuthServer(authContext)

	go as.startAuthenticationServer(authContext)
	as.authState = oauth2.GenerateVerifier()

	if err := as.requestAuthCode(); err != nil {
		return err
	}

	as.Token = <-as.authorizationCompleteChan
	cancel()
	return nil
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

	mux.Handle("/complete", as.AuthCompleteHandler())

	as.authSever = &http.Server{
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
		err := as.authSever.ListenAndServe()
		utils.Check(err)
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := as.authSever.Shutdown(shutdownCtx)
	utils.Check(err)
}
