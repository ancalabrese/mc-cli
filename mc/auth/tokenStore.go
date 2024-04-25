package auth

import (
	"github.com/hashicorp/go-hclog"
	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
)

type TokenStore interface {
	Get(clientId string) (*oauth2.Token, error)
	Save(clientId string, t *oauth2.Token) error
}

// Implements TokenStore. Stores and retrieve the token from
// keyring
type KeyringTokenStore struct {
	l hclog.Logger
}

func (kts *KeyringTokenStore) Get(clientId string) (*oauth2.Token, error) {
	rt, err := keyring.Get("mc-cli", clientId)
	if err != nil {
		kts.l.Error("KeyringStore", "failed to save access token", err)
		return nil, err
	}

	return &oauth2.Token{
		RefreshToken: rt,
	}, nil
}

func (kts *KeyringTokenStore) Save(clientId string, t *oauth2.Token) error {
	err := keyring.Set("mc-cli", clientId, t.RefreshToken)
	kts.l.Error("KeyringStore", "failed to get access token", err)
	return err
}
