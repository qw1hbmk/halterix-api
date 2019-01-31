package platform

import (
	//"google.golang.org/api/iterator"
	"os"
    "path/filepath"
	// "net/http"

	"cloud.google.com/go/firestore"

	//"github.com/julienschmidt/httprouter"
	"context"
	"log"

	firebase "firebase.google.com/go"
	//"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFireStoreConnection(credsFile string, ctx context.Context) *firestore.Client {

	// Get path of local resource file if not provided

	if credsFile == "" {
		cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		trimmedPath := cwd[:len(cwd)-3]
		credsFile =  trimmedPath + "/internal/platform/halterixapi-firebase-adminsdk.json"
	}

	// Set up firestore

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