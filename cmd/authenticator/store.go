package authenticator

import (
	"context"

	"cloud.google.com/go/firestore"
)

type database struct {
	store *firestore.Client
	ctx   context.Context
}

func NewDatabase(store *firestore.Client, ctx context.Context) *database {
	return &database{store, ctx}
}

func (db *database) GetAuthDetails(apiKey string) (*AccountAuth, error) {

	dsnap, err := db.store.Collection("authorized-keys").Doc(apiKey).Get(db.ctx)
	if err != nil {
		return &AccountAuth{}, err
	}
	var a AccountAuth
	err = dsnap.DataTo(&a)
	if err != nil {
		return &AccountAuth{}, err
	}
	a.APIKey = apiKey
	return &a, nil
}
