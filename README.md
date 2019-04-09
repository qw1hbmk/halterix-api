# halterix-api

### Code Details

Since the functionality of this API is quite small, all the endpoints are currently listed in one package: watchtower.

The program begins in main.go.

For details on: 
- which routes (endpoints) can be hit, go to watchtower/routes.go
- the handlers (behaviour) of each route, go to watchtower/handlers.go
- interactions with firestore, go to watchtower/store.go

The dependency manager used is glide. If using new external packages, import them within your file, and run 

``glide update``

Be sure to commit the glide.yaml and glide.lock files.

### Test Locally

Run: 

``go build``

``./haterix-api``

This will launch the program on localhost:8080, and it will be configured with the database set in main.go (currently, haterix-dev).
If you have a permissions error accessing the database, download the services json file, and run: 

``export GOOGLE_APPLICATION_CREDENTIALS="<location of file>/halterixadmin-sdk.json"``


### Deployment

To deploy new versions of the application, run: 

``gcloud app deploy --project <project-id>``

The three environments we can deploy to are: halterix-dev, halterix-rnd, and halerix-prod.

