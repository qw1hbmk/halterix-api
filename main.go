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

	router := httprouter.New()

	router.GET("/", homeHandler)

	ctx := context.Background()

	// TO DO: Pass in config file from firetore to increase security
	firestore := platform.NewFireStoreConnection("", ctx)

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
