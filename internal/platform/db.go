package platform

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func NewFireStoreConnection(credsFile string, ctx context.Context) *firestore.Client {

	conf := &firebase.Config{ProjectID: "halterixapi"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
