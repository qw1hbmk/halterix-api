package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/qw1hbmk/halterix-api/cmd/authenticator"
	"github.com/qw1hbmk/halterix-api/cmd/watchtower"
	"github.com/qw1hbmk/halterix-api/internal/platform"

	"github.com/julienschmidt/httprouter"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {

	// var1 := os.Getenv("MY_VAR")
	// fmt.Println(var1)

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	// Todo: This should be set in a bash script and passed as
	// an env variable
	var fireStoreId string
	cloudId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if cloudId == "halterix-api-prod" || cloudId == "halterix-prod-83363" {
		fireStoreId = "halterix-prod-83363"
	} else if cloudId == "halterix-api-rnd" || cloudId == "spars-9-axis" {
		fireStoreId = "spars-9-axis"
	} else if cloudId == "halterix-api-dev" || cloudId == "halterix-dev" {
		fireStoreId = "halterix-dev"
	} else {
		// Localhost: still set firestore to dev, but set logging locally
		fireStoreId = "halterix-dev"

		// Set Logger
		// var filename string = "logfile.log"
		// f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
		// log.SetOutput(fi)
	}

	// Set up firestore
	ctx := context.Background()
	firestore := platform.NewFireStoreConnection(fireStoreId, ctx)

	router := httprouter.New()
	router.GET("/", homeHandler)

	// Apply auth middleware
	adb := authenticator.NewDatabase(firestore, ctx)
	authedRouter := authenticator.APIKeyMiddleware(router, adb)

	// Set up watchtower endpoints
	wdb := watchtower.NewDatabase(firestore, ctx)
	ws := watchtower.NewServer(router, wdb)

	ws.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	logrus.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), authedRouter))
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome to " + os.Getenv("GOOGLE_CLOUD_PROJECT")))
}
