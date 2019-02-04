package platform

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFireStoreConnection(credsFile string, ctx context.Context) *firestore.Client {

	// Get path of local resource file if not provided
	if credsFile == "" {
		credsFile = "creds/halterixapi-firebase-adminsdk.json"
	}

	opt := option.WithCredentialsFile(credsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
