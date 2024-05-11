package storage

import (
	"github.com/hashicorp/go-hclog"
	"github.com/zalando/go-keyring"
)

const KEYRING_SERVICE = "mc-cli"

type SecretStore interface {
	Get(key string) (string, error)
	Save(key string, v string) error
}

// Implements TokenStore. Stores and retrieve the token from
// keyring
type KeyringStore struct {
	l hclog.Logger
}

func NewKeyRingStore(l hclog.Logger) SecretStore {
	return &KeyringStore{l: l.Named("KeyringStore")}
}

func (ks *KeyringStore) Get(key string) (string, error) {
	s, err := keyring.Get(KEYRING_SERVICE, key)
	if err != nil {
		ks.l.Error("KeyringStore", "failed to save secret", err)
		return "", err
	}

	return s, nil
}

func (ks *KeyringStore) Save(key string, v string) error {
	err := keyring.Set(KEYRING_SERVICE, key, v)
	if err != nil {
		ks.l.Error("KeyringStore", "failed to retrieve secret", err)
		return err
	}
	ks.l.Debug("Write success", "key", key, "value", v)
	return nil
}

type ApiSecretStore struct {
	store SecretStore
}

func NewApiSecretStore(l hclog.Logger) *ApiSecretStore {
	return &ApiSecretStore{
		store: NewKeyRingStore(l),
	}
}

func (ass *ApiSecretStore) SaveApiSecret(clientId string, secret string) error {
	return ass.store.Save(clientId, secret)
}

func (ass *ApiSecretStore) SaveRefreshAccessToken(clientId string, rt string) error {
	return ass.store.Save("r_token_"+clientId, rt)
}

func (ass *ApiSecretStore) GetApiSecret(clientId string) (string, error) {
	return ass.store.Get(clientId)
}

func (ass *ApiSecretStore) GetRefreshAccessToken(clientId string) (string, error) {
	return ass.store.Get("r_token_" + clientId)
}
