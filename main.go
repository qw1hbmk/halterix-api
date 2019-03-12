package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qw1hbmk/halterix-api/cmd/watchtower"
	"github.com/qw1hbmk/halterix-api/internal/platform"

	"github.com/julienschmidt/httprouter"
)

func main() {

	var1 := os.Getenv("MY_VAR")
	fmt.Println(var1)

	// Todo: This should be set in a bash script and passed as
	// an env variable
	var fireStoreId string
	cloudId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if cloudId == "halterix-api-prod" {
		fireStoreId = "halterix-prod"
	} else if cloudId == "halterix-api-rnd" {
		fireStoreId = "spars-9-axis"
	} else {
		// All remaining projects should go to the dev database
		fireStoreId = "halterix-dev"
	}

	router := httprouter.New()
	router.GET("/", homeHandler)

	ctx := context.Background()
	firestore := platform.NewFireStoreConnection(fireStoreId, ctx)

	// Set up watchtower endpoints
	wdb := watchtower.NewDatabase(firestore, ctx)
	ws := watchtower.NewServer(router, wdb)
	ws.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome"))
}
