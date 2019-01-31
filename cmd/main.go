package main

import (
	//"google.golang.org/api/iterator"
	"fmt"
	 "net/http"

	//"cloud.google.com/go/firestore"

    "context"
    //"log"

    //firebase "firebase.google.com/go"
    //"firebase.google.com/go/auth"
	//"google.golang.org/api/option"
	
	"github.com/julienschmidt/httprouter"

	"github.com/qw1hbmk/halterix-api/internal/platform"
	"github.com/qw1hbmk/halterix-api/cmd/watchtower"
)

func main() {

	router := httprouter.New()

	ctx := context.Background()

	// TO DO: Pass in config file from firetore to increase security
	firestore := platform.NewFireStoreConnection("", ctx)

	// Set up watchtower endpoints
	wdb := watchtower.NewDatabase(firestore, ctx)
	ws := watchtower.NewServer(router, wdb)
	ws.RegisterRoutes()

	// Listen
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		fmt.Println(err)
	}	

}